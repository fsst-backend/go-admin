package name

import (
	"context"
	"errors"

	"go-admin/common/errcode"
	"go-admin/grpc/pb"

	"github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	ext "go-admin/config"
)

type AdminService struct {
	pb.UnimplementedAdminUserServiceServer
	Logger *logger.Helper
	Orm    *gorm.DB
	Config *ext.Extend // 扩展配置（环境变量）
}

// GetOrm 获取Orm DB
func (s *AdminService) GetOrm() (*gorm.DB, error) {
	// 直接从 Runtime 获取默认数据库
	dbMap := sdk.Runtime.GetDb()
	if len(dbMap) == 0 {
		return nil, errors.New("数据库连接未初始化")
	}
	// 获取第一个数据库连接
	for _, db := range dbMap {
		return db, nil
	}
	return nil, errors.New("无法获取数据库连接")
}

// MakeOrm 设置Orm DB
func (s *AdminService) MakeOrm() *AdminService {
	// 直接从 Runtime 获取默认数据库
	dbMap := sdk.Runtime.GetDb()
	if len(dbMap) == 0 {
		if s.Logger != nil {
			s.Logger.Error(500, errors.New("数据库连接未初始化"), "数据库连接获取失败")
		}
		return s
	}
	// 获取第一个数据库连接
	for _, db := range dbMap {
		s.Orm = db
		break
	}
	return s
}

func (s *AdminService) GetAdminUserInfo(ctx context.Context, req *pb.GetAdminUserReq) (*pb.GetAdminUserResp, error) {
	// 确保 ORM 已初始化
	if s.Orm == nil {
		var err error
		s.Orm, err = s.GetOrm()
		if err != nil {
			if s.Logger != nil {
				s.Logger.Error(500, err, "获取数据库连接失败")
			}
			return &pb.GetAdminUserResp{
				Status: &pb.Status{
					Code:    int32(errcode.RespCodeAccountServerError),
					Message: "数据库连接失败: " + err.Error(),
				},
			}, err
		}
	}

	// 通过 UUID 查询用户
	var user models.SysUser
	err := s.Orm.Where("uuid = ?", req.Uid).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if s.Logger != nil {
				s.Logger.Warnf("用户未找到, UUID: %s", req.Uid)
			}
			return &pb.GetAdminUserResp{
				Status: &pb.Status{
					Code:    int32(errcode.RespCodeBadRequest),
					Message: "user not found",
				},
			}, nil
		}
		if s.Logger != nil {
			s.Logger.Error(500, err, "查询用户失败")
		}
		return &pb.GetAdminUserResp{
			Status: &pb.Status{
				Code:    int32(errcode.RespCodeAccountServerError),
				Message: "查询用户失败: " + err.Error(),
			},
		}, err
	}

	// 检查用户是否存在
	if user.UserId == 0 {
		return &pb.GetAdminUserResp{
			Status: &pb.Status{
				Code:    int32(errcode.RespCodeBadRequest),
				Message: "user not found",
			},
		}, nil
	}

	return &pb.GetAdminUserResp{
		Status: &pb.Status{
			Code: 0,
		},
		Data: &pb.AdminUser{
			Id:     int32(user.UserId),
			Uid:    user.UUID,
			Name:   user.Username,
			Alias:  user.NickName,
			Email:  user.Email,
			Avatar: user.Avatar,
			Tel:    user.Phone,
			State:  int32(user.Status), // Status 已经是 int 类型，直接转换为 int32
		},
	}, nil
}

func (s *AdminService) GetAdminUserList(ctx context.Context, req *pb.GetAdminUserListReq) (*pb.GetAdminUserListResp, error) {
	// 确保 ORM 已初始化
	if s.Orm == nil {
		var err error
		s.Orm, err = s.GetOrm()
		if err != nil {
			if s.Logger != nil {
				s.Logger.Error(500, err, "获取数据库连接失败")
			}
			return &pb.GetAdminUserListResp{
				Status: &pb.Status{
					Code:    int32(errcode.RespCodeAccountServerError),
					Message: "数据库连接失败: " + err.Error(),
				},
			}, err
		}
	}

	// UUID 列表去重
	req.Uids = uniqStrs(req.Uids)
	if len(req.Uids) == 0 {
		return &pb.GetAdminUserListResp{
			Status: &pb.Status{
				Success: true,
				Code:    0,
				Message: "success",
			},
			Data: []*pb.AdminUser{},
		}, nil
	}

	// 通过 UUID 列表查询用户
	var users []models.SysUser
	err := s.Orm.Where("uuid IN ?", req.Uids).Find(&users).Error
	if err != nil {
		if s.Logger != nil {
			s.Logger.Error(500, err, "查询用户列表失败")
		}
		return &pb.GetAdminUserListResp{
			Status: &pb.Status{
				Code:    int32(errcode.RespCodeAccountServerError),
				Message: "查询用户列表失败: " + err.Error(),
			},
		}, err
	}

	return &pb.GetAdminUserListResp{
		Status: &pb.Status{
			Success: true,
			Code:    0,
			Message: "success",
		},
		Data: convertSysUsersToRpcUsers(users),
	}, nil
}

// uniqStrs 字符串列表去重
func uniqStrs(strs []string) []string {
	if len(strs) == 0 {
		return []string{}
	}
	seen := make(map[string]struct{})
	result := make([]string, 0, len(strs))
	for _, s := range strs {
		if s == "" {
			continue
		}
		if _, exists := seen[s]; !exists {
			seen[s] = struct{}{}
			result = append(result, s)
		}
	}
	return result
}

// convertSysUsersToRpcUsers 将 SysUser 列表转换为 RPC AdminUser 列表
func convertSysUsersToRpcUsers(users []models.SysUser) []*pb.AdminUser {
	rpcUsers := make([]*pb.AdminUser, 0, len(users))
	for _, user := range users {
		rpcUsers = append(rpcUsers, &pb.AdminUser{
			Id:     int32(user.UserId),
			Uid:    user.UUID,
			Name:   user.Username,
			Alias:  user.NickName,
			Email:  user.Email,
			Avatar: user.Avatar,
			Tel:    user.Phone,
			State:  int32(user.Status),
		})
	}
	return rpcUsers
}

func (s *AdminService) GetAllDepartments(ctx context.Context, in *pb.GetAllDepartmentReq) (*pb.GetAllDepartmentResppone, error) {
	// 确保 ORM 已初始化
	if s.Orm == nil {
		var err error
		s.Orm, err = s.GetOrm()
		if err != nil {
			if s.Logger != nil {
				s.Logger.Error(500, err, "获取数据库连接失败")
			}
			return &pb.GetAllDepartmentResppone{
				Status: &pb.Status{
					Code:    int32(errcode.RespCodeAccountServerError),
					Message: "数据库连接失败: " + err.Error(),
				},
			}, err
		}
	}

	// 查询部门数据（根据 catalog 过滤）
	var deptList []models.SysDept
	query := s.Orm.Where("status = ?", 1) // 只查询启用状态的部门
	if in.Catalog != "" {
		query = query.Where("dept_catalog = ?", in.Catalog)
	}
	err := query.Order("sort ASC, dept_id ASC").Find(&deptList).Error
	if err != nil {
		if s.Logger != nil {
			s.Logger.Error(500, err, "查询部门列表失败")
		}
		return &pb.GetAllDepartmentResppone{
			Status: &pb.Status{
				Code:    int32(errcode.RespCodeAccountServerError),
				Message: "查询部门列表失败: " + err.Error(),
			},
		}, err
	}

	if len(deptList) == 0 {
		return &pb.GetAllDepartmentResppone{
			Status: &pb.Status{
				Success: true,
				Code:    0,
				Message: "success",
			},
			Data: []*pb.Department{},
		}, nil
	}

	// 构建部门树形结构
	deptTree := buildDeptTree(deptList)

	// 收集所有部门ID
	deptIds := getAllDeptIds(deptTree)

	// 获取负责人和成员
	leadersM, membersM, err := s.getLeadersAndMembers(deptIds, deptList)
	if err != nil {
		if s.Logger != nil {
			s.Logger.Error(500, err, "获取部门负责人和成员失败")
		}
		return &pb.GetAllDepartmentResppone{
			Status: &pb.Status{
				Code:    int32(errcode.RespCodeAccountServerError),
				Message: "获取部门负责人和成员失败: " + err.Error(),
			},
		}, err
	}

	// 转换为 RPC 格式
	rpcDepts := convertDeptTreeToRpcDepts(deptTree, leadersM, membersM)

	return &pb.GetAllDepartmentResppone{
		Status: &pb.Status{
			Success: true,
			Code:    0,
			Message: "success",
		},
		Data: rpcDepts,
	}, nil
}

// DeptTreeNode 部门树节点
type DeptTreeNode struct {
	Dept     models.SysDept
	Children []*DeptTreeNode
}

// buildDeptTree 构建部门树形结构
func buildDeptTree(deptList []models.SysDept) []*DeptTreeNode {
	// 创建部门映射
	deptMap := make(map[int]*DeptTreeNode)
	for i := range deptList {
		deptMap[deptList[i].DeptId] = &DeptTreeNode{
			Dept:     deptList[i],
			Children: []*DeptTreeNode{},
		}
	}

	// 构建树形结构
	var rootNodes []*DeptTreeNode
	for i := range deptList {
		node := deptMap[deptList[i].DeptId]
		if deptList[i].ParentId == 0 {
			rootNodes = append(rootNodes, node)
		} else {
			if parent, ok := deptMap[deptList[i].ParentId]; ok {
				parent.Children = append(parent.Children, node)
			}
		}
	}

	return rootNodes
}

// getAllDeptIds 获取所有部门ID（包括子部门）
func getAllDeptIds(nodes []*DeptTreeNode) []int {
	var ids []int
	for _, node := range nodes {
		ids = append(ids, node.Dept.DeptId)
		ids = append(ids, getAllDeptIds(node.Children)...)
	}
	return ids
}

// getLeadersAndMembers 获取部门的负责人和成员
func (s *AdminService) getLeadersAndMembers(deptIds []int, deptList []models.SysDept) (map[int][]models.SysUser, map[int][]models.SysUser, error) {
	leadersM := make(map[int][]models.SysUser)
	membersM := make(map[int][]models.SysUser)

	if len(deptIds) == 0 {
		return leadersM, membersM, nil
	}

	// 收集所有负责人的 UUID
	leaderUUIDs := make([]string, 0)
	deptLeaderMap := make(map[string][]int) // UUID -> []DeptId
	for _, dept := range deptList {
		if dept.Leader != "" {
			leaderUUIDs = append(leaderUUIDs, dept.Leader)
			deptLeaderMap[dept.Leader] = append(deptLeaderMap[dept.Leader], dept.DeptId)
		}
	}

	// 查询负责人信息
	if len(leaderUUIDs) > 0 {
		var leaders []models.SysUser
		err := s.Orm.Where("uuid IN ?", leaderUUIDs).Find(&leaders).Error
		if err != nil {
			return nil, nil, err
		}
		// 将负责人分配到对应部门
		for _, leader := range leaders {
			if deptIds, ok := deptLeaderMap[leader.UUID]; ok {
				for _, deptId := range deptIds {
					leadersM[deptId] = append(leadersM[deptId], leader)
				}
			}
		}
	}

	// 查询部门成员（通过 dept_id 关联）
	var members []models.SysUser
	err := s.Orm.Where("dept_id IN ?", deptIds).Find(&members).Error
	if err != nil {
		return nil, nil, err
	}
	// 将成员分配到对应部门
	for _, member := range members {
		if member.DeptId > 0 {
			membersM[member.DeptId] = append(membersM[member.DeptId], member)
		}
	}

	return leadersM, membersM, nil
}

// convertDeptTreeToRpcDepts 将部门树转换为 RPC 格式
func convertDeptTreeToRpcDepts(nodes []*DeptTreeNode, leadersM, membersM map[int][]models.SysUser) []*pb.Department {
	rpcDepts := make([]*pb.Department, 0, len(nodes))
	for _, node := range nodes {
		rpcDept := &pb.Department{
			Id:      int32(node.Dept.DeptId),
			Alias:   node.Dept.DeptName,
			Brief:   "", // SysDept 没有 Brief 字段
			LogoUrl: "", // SysDept 没有 LogoURL 字段
			Catalog: node.Dept.DeptCatalog,
			Leader:  convertSysUsersToRpcUsers(leadersM[node.Dept.DeptId]),
			Member:  convertSysUsersToRpcUsers(membersM[node.Dept.DeptId]),
		}
		// 递归处理子部门
		if len(node.Children) > 0 {
			rpcDept.Children = convertDeptTreeToRpcDepts(node.Children, leadersM, membersM)
		}
		rpcDepts = append(rpcDepts, rpcDept)
	}
	return rpcDepts
}
