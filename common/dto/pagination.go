package dto

type Pagination struct {
	PageIndex int `form:"pageIndex"`
	PageSize  int `form:"pageSize"`
}

func (m *Pagination) GetPageIndex() int {
	if m.PageIndex <= 0 {
		m.PageIndex = 1
	}
	return m.PageIndex
}

func (m *Pagination) GetPageSize() int {
	if m.PageSize <= 0 {
		m.PageSize = 10
	}
	return m.PageSize
}

// OffsetLimitPagination 使用 offset 和 limit 的分页结构
type OffsetLimitPagination struct {
	Limit  int `form:"limit"`
	Offset int `form:"offset"`
}

// GetLimit 获取限制数量，如果小于等于0则返回默认值10
func (o *OffsetLimitPagination) GetLimit() int {
	if o.Limit <= 0 {
		return 10
	}
	return o.Limit
}

// GetOffset 获取偏移量，如果小于0则返回0
func (o *OffsetLimitPagination) GetOffset() int {
	if o.Offset < 0 {
		return 0
	}
	return o.Offset
}

// GetOffsetLimitPage 从PageIndex/PageSize转换为Offset/Limit
func (m *Pagination) GetOffsetLimitPage() (limit, offset int) {
	pageIndex := m.GetPageIndex()
	pageSize := m.GetPageSize()
	offset = (pageIndex - 1) * pageSize
	if offset < 0 {
		offset = 0
	}
	limit = pageSize
	return limit, offset
}
