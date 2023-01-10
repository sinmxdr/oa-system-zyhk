package magic

const (

	//合同人民币回款
	CONTRACT_PAY_TYPE_CNY = 1
	//合同美元回款
	CONTRACT_PAY_TYPE_USD = 2

	//审批驳回
	CONTRACT_STATUS_REJECT = -1
	//保存
	CONTRACT_STATUS_SAVE = 0
	//未审批
	CONTRACT_STATUS_NOT_APPROVAL = 1
	//未完成
	CONTRACT_STATUS_NOT_FINISH = 2
	//已完成
	CONTRACT_STATUS_FINISH = 3

	//生产中
	CONTRATCT_PRODUCTION_STATUS_ING = 1
	//生产完成
	CONTRATCT_PRODUCTION_STATUS_FINISH = 2

	//回款中
	CONTRATCT_COLLECTION_STATUS_ING = 1
	//回款完成
	CONTRATCT_COLLECTION_STATUS_FINISH = 2

	//合同提成系统结算
	CONTRATCT_PUSHMONEY_AUTO = 1
	//合同提成手动结算
	CONTRATCT_PUSHMONEY_MANUAL = 2

	//驳回
	TASK_STATUS_REJECT = -1
	//待分发
	TASK_STATUS_NOT_DISTRIBUTE = 0
	//待设计
	TASK_STATUS_NOT_DESIGN = 1
	//待采购
	TASK_STATUS_NOT_PURCHASE = 2
	//待入/出库
	TASK_STATUS_NOT_STORAGE = 3
	//待装配（非标设计流程独占）
	TASK_STATUS_NOT_ASSEMBLY = 4
	//待发货
	TASK_STATUS_NOT_SHIPMENT = 5
	//已发货
	TASK_STATUS_SHIPMENT = 6

	//标准/第三方有库存
	TASK_TYPE_1 = 1
	//标准/第三方无库存
	TASK_TYPE_2 = 2
	//非标准定制
	TASK_TYPE_3 = 3

	// 财务审批状态
	EXPENSE_STATUS_FAIL           = -1
	EXPENSE_STATUS_NOT_APPROVAL_1 = 1
	EXPENSE_STATUS_NOT_APPROVAL_2 = 2
	EXPENSE_STATUS_NOT_PAYMENT    = 3
	EXPENSE_STATUS_FINISH         = 4

	// 财务类型
	EXPENSE_TYPE_BuZhu    = 1
	EXPENSE_TYPE_TiChen   = 2
	EXPENSE_TYPE_YeWuFei  = 3
	EXPENSE_TYPE_CaiLvFei = 4

	//保证金状态
	BIDBOND_STATUS_FAIL         = -1
	BIDBOND_STATUS_NOT_APPROVAL = 1
	BIDBOND_STATUS_NOT_FINAL    = 2
	BIDBOND_STATUS_FINAL        = 3

	//预设计状态
	PREDESIGN_STATUS_FAIL         = -1
	PREDESIGN_STATUS_NOT_APPROVAL = 1
	PREDESIGN_STATUS_NOT_FINAL    = 2
	PREDESIGN_STATUS_FINAL        = 3

	//预设计任务状态
	PREDESIGN_TASK_STATUS_NOT_SUBMIT   = 1
	PREDESIGN_TASK_STATUS_NOT_APPROVAL = 2
	PREDESIGN_TASK_STATUS_FAIL         = 3
	PREDESIGN_TASK_STATUS_FINAL        = 4

	//产品类型分类
	PRODUCT_TYPE_TYPE_YW = 1
	PRODUCT_TYPE_TYPE_ZY = 1
	PRODUCT_TYPE_TYPE_QD = 1

	//采购状态
	PURCHASING_STATUS_REJECT     = -1
	PURCHASING_STATUS_SAVE       = 0
	PURCHASING_STATUS_NO_CHECK   = 1
	PURCHASING_STATUS_NO_APPROVE = 2
	PURCHASING_STATUS_NO_FINAL   = 3
	PURCHASING_STATUS_FINAL      = 4

	PURCHASING_PRODUCT_STATUS_NO_FINAL = 1
	PURCHASING_PRODUCT_STATUS_FINAL    = 2

	PURCHASING_PAY_STATUS_NO_FINAL = 1
	PURCHASING_PAY_STATUS_FINAL    = 2

	PURCHASING_INVOICE_STATUS_NO_FINAL = 1
	PURCHASING_INVOICE_STATUS_FINAL    = 2

	PURCHASING_TYPE_CONTRACT = 1
	PURCHASING_TYPE_TASK     = 2
	PURCHASING_TYPE_BANK     = 3
)
