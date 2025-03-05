package response

// Success
type TResponseMeta struct {
	Status int    `json:"status"`
	Remark string `json:"remark"` 
}

type TSuccessResponse struct {
	Meta    TResponseMeta `json:"meta"`
	Results interface{}   `json:"results"`
}

func SuccessResponse(status int, remark string, data interface{}) interface{} {
	if data == nil {
		return TErrorResponse{
			Meta: TResponseMeta{
				Status: status,
				Remark: remark,
			},
		}
	} else {
		return TSuccessResponse{
			Meta: TResponseMeta{
				Status: status,
				Remark: remark,
			},
			Results: data,
		}
	}
}

// Error
type TErrorResponse struct {
	Meta TResponseMeta `json:"meta"`
}

func ErrorResponse(status int, remark string) interface{} {
	return TErrorResponse{
		Meta: TResponseMeta{
			Status: status,
			Remark: remark,
		},
	}
}

// Pagination
type TResponseMetaPage struct {
	Status         int    `json:"status"`
	Remark         string `json:"remark"` 
	CurrentPage    int    `json:"current_page"`
	TotalPages     int    `json:"total_pages"`
	TotalItems     int    `json:"total_items"`
	ItemsPerPage   int    `json:"items_per_page"`
	HasNextPage    bool   `json:"has_next_page"`
	HasPreviousPage bool  `json:"has_previous_page"`
}

type TSuccessResponsePage struct {
	Meta    TResponseMetaPage `json:"meta"`
	Results interface{}       `json:"results"`
}

func SuccessResponsePage(status int, remark string, page int, limit int, totaldata int64, data interface{}) TSuccessResponsePage {
	totalPages := int(totaldata / int64(limit))
	if totaldata % int64(limit) != 0 {
		totalPages++
	}

	return TSuccessResponsePage{
		Meta: TResponseMetaPage{
			Status:         status,
			Remark:         remark,
			CurrentPage:    page,
			TotalPages:     totalPages,
			TotalItems:     int(totaldata),
			ItemsPerPage:   limit,
			HasNextPage:    page < totalPages,
			HasPreviousPage: page > 1,
		},
		Results: data,
	}
}
