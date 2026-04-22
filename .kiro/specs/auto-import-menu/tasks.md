# 实现计划：启动时自动导入菜单配置

## 概述

将菜单配置文件移至 `config/` 目录，修改 Dockerfile 添加 COPY 指令，在 `cmd/api/server.go` 中实现 `importMenuFromConfig` 函数，并在 `setup()` 中 SysApi Swagger 同步之后调用。所有错误记录日志但不阻塞启动。属性测试使用 `rapid` 库验证 JSON 序列化往返一致性和结构验证正确性。

## 任务

- [x] 1. 移动菜单配置文件到 config/ 目录
  - 将项目根目录的 `menu_20260416180529.json` 移动到 `config/menu_20260416180529.json`
  - 确认文件内容为有效 JSON，顶层包含 `menus` 数组和 `permissions` 数组
  - _需求：1.1, 1.2_

- [x] 2. 修改 Dockerfile 添加 COPY 指令
  - 在 `Dockerfile` 现有 COPY 指令区域添加一行：`COPY config/menu_20260416180529.json ./menu_20260416180529.json`
  - 放置在 `COPY config/settings.template.yaml ./settings.yaml` 之前或之后均可，保持与其他 config COPY 指令风格一致
  - _需求：2.1, 2.2_

- [x] 3. 实现 importMenuFromConfig 函数并集成到 setup()
  - [x] 3.1 在 `cmd/api/server.go` 中实现 `importMenuFromConfig(db *gorm.DB, lg *logger.Helper)` 函数
    - 检查 `db` 是否为 nil，若为 nil 则 `lg.Warn("启动时菜单导入跳过：数据库连接为空")` 并返回
    - 使用 `os.ReadFile("menu_20260416180529.json")` 读取配置文件
    - 若文件不存在（`os.IsNotExist(err)`），记录 Warn 日志并返回
    - 若为其他读取错误，记录 Error 日志并返回
    - 使用 `json.Unmarshal` 反序列化为 `dto.MenuPermissionIO`
    - 若 JSON 解析失败，记录 Error 日志并返回
    - 初始化 `service.SysMenu`，设置 `svc.Orm = db` 和 `svc.Log = lg`
    - 调用 `svc.ImportMenuPermission(&data)`
    - 若导入失败，记录 Error 日志并返回
    - 导入成功，记录 Info 日志
    - 函数不返回 error，所有错误场景均记录日志后直接 return
    - _需求：3.1, 3.2, 3.3, 4.1, 4.2, 4.3, 4.4_

  - [x] 3.2 在 `setup()` 函数中调用 `importMenuFromConfig`
    - 在 SysApi Swagger 同步代码块之后、注册监听函数之前插入调用
    - 复用已有的 `db` 和 `h` 变量：`importMenuFromConfig(db, h)`
    - 需要在 `setup()` 中添加 `encoding/json` 和 `go-admin/app/admin/service` 的 import（如果尚未存在）
    - _需求：3.4_

- [x] 4. 检查点 - 确保编译通过
  - 确保所有代码编译通过，ask the user if questions arise.

- [x] 5. 属性测试与单元测试
  - [x] 5.1 编写属性测试：MenuPermissionIO JSON 序列化往返
    - **属性 1：MenuPermissionIO JSON 序列化往返**
    - **验证需求：3.1**
    - 在 `cmd/api/server_test.go`（或 `app/admin/service/dto/sys_menu_permission_import_test.go`）中使用 `rapid` 库
    - 使用 `rapid.Custom` 生成随机的 `dto.MenuPermissionIO` 结构体（随机数量的菜单和权限，菜单支持随机深度子菜单树）
    - 序列化为 JSON 后反序列化，使用 `reflect.DeepEqual` 验证结果与原始值一致
    - 最少 100 次迭代
    - 注释标注：`Feature: auto-import-menu, Property 1: MenuPermissionIO JSON 序列化往返`

  - [x] 5.2 编写属性测试：JSON 结构验证正确性
    - **属性 2：JSON 结构验证正确性**
    - **验证需求：1.2, 3.1**
    - 使用 `rapid` 生成随机 JSON 字节流（包括有效结构、缺少字段、类型错误等）
    - 验证有效结构能正确解析，无效结构不会导致 panic 且结果为零值
    - 最少 100 次迭代
    - 注释标注：`Feature: auto-import-menu, Property 2: JSON 结构验证正确性`

  - [x] 5.3 编写单元测试：importMenuFromConfig 错误处理场景
    - 测试文件不存在场景：传入不存在的文件路径，验证函数正常返回不 panic
    - 测试无效 JSON 场景：创建包含无效 JSON 的临时文件，验证函数记录错误日志并正常返回
    - 测试数据库连接为 nil 场景：传入 nil db，验证函数记录警告并正常返回
    - _需求：4.1, 4.2, 4.4_

- [x] 6. 最终检查点 - 确保所有测试通过
  - 确保所有测试通过，ask the user if questions arise.

## 备注

- 标记 `*` 的任务为可选任务，可跳过以加速 MVP 交付
- 每个任务引用了具体的需求编号以确保可追溯性
- 检查点确保增量验证
- 属性测试使用 `rapid` 库（Go 语言主流 PBT 库）验证正确性属性
- 单元测试验证具体的错误处理场景和边界条件
- `importMenuFromConfig` 函数不返回 error，与设计文档一致：所有错误场景行为一致（记录日志并继续）
