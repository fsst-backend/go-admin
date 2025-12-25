#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
从 Swagger JSON 文件生成 sys_api 表的 SQL 插入语句
"""

import json
import os
import sys

def generate_api_sql(swagger_file, output_file=None):
    """
    从 swagger.json 生成 sys_api 表的 SQL 插入语句
    
    Args:
        swagger_file: swagger.json 文件路径
        output_file: 输出 SQL 文件路径，如果为 None 则输出到标准输出
    """
    # 读取 swagger.json
    with open(swagger_file, 'r', encoding='utf-8') as f:
        swagger = json.load(f)
    
    # 收集所有 API
    apis = []
    for path, methods in swagger.get('paths', {}).items():
        for method, details in methods.items():
            title = details.get('summary', '未知')
            tags = details.get('tags', ['其他'])
            tag = tags[0] if tags else '其他'
            
            # 确定接口类型
            # SYS: 系统管理相关接口
            # BUS: 业务接口
            api_type = 'SYS' if path.startswith('/lotus/api/v1/sys-') else 'BUS'
            
            apis.append({
                'path': path,
                'method': method.upper(),
                'title': title,
                'type': api_type,
                'tag': tag
            })
    
    # 按路径和方法排序
    apis.sort(key=lambda x: (x['path'], x['method']))
    
    # 生成 SQL
    lines = []
    lines.append("-- ============================================")
    lines.append("-- sys_api 表数据插入脚本")
    lines.append("-- 从 Swagger 文档自动生成")
    lines.append("-- ============================================")
    lines.append("")
    lines.append("-- 表结构:")
    lines.append("-- id: 主键ID")
    lines.append("-- handle: 处理器名称(暂时为空)")
    lines.append("-- title: API标题/描述")
    lines.append("-- path: API路径")
    lines.append("-- action: HTTP方法(GET/POST/PUT/DELETE)")
    lines.append("-- type: 接口类型(SYS=系统/BUS=业务)")
    lines.append("-- created_at: 创建时间")
    lines.append("-- updated_at: 更新时间")
    lines.append("-- deleted_at: 删除时间(软删除)")
    lines.append("-- create_by: 创建者ID")
    lines.append("-- update_by: 更新者ID")
    lines.append("")
    
    for idx, api in enumerate(apis, start=1):
        # 转义单引号
        title = api['title'].replace("'", "''")
        
        sql = (
            f"INSERT INTO sys_api VALUES "
            f"({idx}, '', '{title}', '{api['path']}', '{api['method']}', "
            f"'{api['type']}', NOW(), NOW(), NULL, 1, 1);"
        )
        lines.append(sql)
    
    lines.append("")
    lines.append(f"-- 总计: {len(apis)} 条 API 记录")
    lines.append("")
    
    # 添加按标签分组的统计
    tag_stats = {}
    for api in apis:
        tag = api['tag']
        tag_stats[tag] = tag_stats.get(tag, 0) + 1
    
    lines.append("-- 按标签分组统计:")
    for tag, count in sorted(tag_stats.items()):
        lines.append(f"-- {tag}: {count} 个接口")
    
    # 输出结果
    result = '\n'.join(lines)
    
    if output_file:
        with open(output_file, 'w', encoding='utf-8') as f:
            f.write(result)
        print(f"✅ SQL 文件已生成: {output_file}")
        print(f"📊 总计 {len(apis)} 条 API 记录")
    else:
        print(result)
    
    return len(apis)


if __name__ == '__main__':
    # 获取项目根目录
    script_dir = os.path.dirname(os.path.abspath(__file__))
    project_root = os.path.dirname(script_dir)
    
    # 默认输入和输出文件
    swagger_file = os.path.join(project_root, 'docs/admin/admin_swagger.json')
    output_file = os.path.join(project_root, 'config/db-api-insert.sql')
    
    # 支持命令行参数
    if len(sys.argv) > 1:
        swagger_file = sys.argv[1]
    if len(sys.argv) > 2:
        output_file = sys.argv[2]
    
    # 检查文件是否存在
    if not os.path.exists(swagger_file):
        print(f"❌ 错误: Swagger 文件不存在: {swagger_file}")
        sys.exit(1)
    
    # 生成 SQL
    try:
        count = generate_api_sql(swagger_file, output_file)
        print(f"✨ 完成!")
    except Exception as e:
        print(f"❌ 错误: {e}")
        import traceback
        traceback.print_exc()
        sys.exit(1)
