# 需求文档

## 简介

本功能实现程序启动时自动导入菜单配置文件（`menu_20260416180529.json`）。该 JSON 文件包含完整的菜单和权限数据，结构与现有 `dto.MenuPermissionIO` 完全匹配。通过在启动流程中自动调用已有的 `SysMenu.ImportMenuPermission` 服务方法，确保每次部署后菜单和权限数据自动同步到数据库，无需手动调用 API。

## 术语表

- **启动器（Startup_Loader）**：程序启动阶段负责读取并导入菜单配置文件的模块，位于 `cmd/api/server.go` 的 `setup()` 函数中
- **菜单配置文件（Menu_Config_File）**：`menu_20260416180529.json` 文件，包含 `menus` 和 `permissions` 两个字段，结构与 `dto.MenuPermissionIO` 一致
- **菜单导入服务（Menu_Import_Service）**：已有的 `service.SysMenu.ImportMenuPermission` 方法，接受 `dto.MenuPermissionIO` 结构体，在事务中先导入权限再导入菜单
- **容器镜像（Container_Image）**：通过 Dockerfile 构建的 Docker 镜像，用于部署运行本系统

## 需求

### 需求 1：菜单配置文件存放

**用户故事：** 作为运维人员，我希望菜单配置文件存放在项目的标准配置目录中，以便统一管理和版本控制。

#### 验收标准

1. THE 菜单配置文件 SHALL 存放在项目的 `config/` 目录下，文件名为 `menu_20260416180529.json`
2. THE 菜单配置文件 SHALL 包含有效的 JSON 内容，且顶层结构包含 `menus` 数组和 `permissions` 数组

### 需求 2：Dockerfile 配置文件复制

**用户故事：** 作为运维人员，我希望 Docker 构建时自动将菜单配置文件复制到容器中，以便容器启动后程序能读取该文件。

#### 验收标准

1. THE Container_Image 的 Dockerfile SHALL 包含一条 COPY 指令，将 `config/menu_20260416180529.json` 复制到容器的 `/app/` 目录下
2. WHEN 容器镜像构建完成后，THE Container_Image SHALL 在 `/app/menu_20260416180529.json` 路径下包含该菜单配置文件

### 需求 3：启动时自动导入菜单数据

**用户故事：** 作为开发人员，我希望程序启动时自动读取菜单配置文件并导入到数据库，以便每次部署后菜单和权限数据自动保持最新。

#### 验收标准

1. WHEN 程序启动且数据库连接可用时，THE Startup_Loader SHALL 读取菜单配置文件并将其内容反序列化为 `MenuPermissionIO` 结构体
2. WHEN 菜单配置文件读取成功后，THE Startup_Loader SHALL 调用 Menu_Import_Service 将菜单和权限数据导入到数据库
3. WHEN 菜单和权限数据导入成功后，THE Startup_Loader SHALL 记录一条包含导入结果的日志信息
4. THE Startup_Loader SHALL 在数据库迁移完成之后、SysApi Swagger 同步之后执行菜单导入操作

### 需求 4：启动时导入错误处理

**用户故事：** 作为开发人员，我希望菜单导入失败时程序能记录错误日志但不阻塞启动，以便其他功能仍可正常运行。

#### 验收标准

1. IF 菜单配置文件不存在，THEN THE Startup_Loader SHALL 记录一条警告日志并跳过菜单导入，程序继续正常启动
2. IF 菜单配置文件内容不是有效的 JSON 格式，THEN THE Startup_Loader SHALL 记录一条错误日志并跳过菜单导入，程序继续正常启动
3. IF Menu_Import_Service 返回错误，THEN THE Startup_Loader SHALL 记录一条包含错误详情的错误日志并跳过菜单导入，程序继续正常启动
4. IF 数据库连接不可用，THEN THE Startup_Loader SHALL 记录一条警告日志并跳过菜单导入
