<template>
    <div>
        <divTable :columnObj="base.column" :tableData="base.tableData" :pageData="base.pageData"
            :handleSizeChange="base.handleSizeChange" :handleCurrentChange="base.handleCurrentChange" />
    </div>
</template>

<script setup>
import { ref, reactive, onBeforeMount } from 'vue'
import divTable from '@/components/divTable/index.vue'

const props = defineProps({
    isRadio: {
        type: Boolean,
        default: false,
    },
    queryObj: {
        type: Object
    },
    queryFunc: {
        type: Function
    },
    arrayObj: {
        type: Array,
        default: []
    },
    chooseFunc: {
        type: Function
    },
    headers: {
        type: Array,
        default: []
    }
})

const base = reactive({
    column: {
        highlightCurrentRow: true,
        headers: [
            {
                type: "operation",
                label: "操作",
                operations: [
                    {
                        isShow: (index, row) => {
                            return props.isRadio
                        },
                        label: "选择",
                        type: "success",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => props.chooseFunc(index, row)
                    },
                    {
                        isShow: (index, row) => {
                            return !props.isRadio && !props.arrayObj.some(item => item.id == row.id)
                        },
                        label: "添加",
                        type: "success",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => base.select(index, row)
                    },
                    {
                        isShow: (index, row) => {
                            return !props.isRadio && props.arrayObj.some(item => item.id == row.id)
                        },
                        label: "移除",
                        type: "danger",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => base.remove(index, row)
                    },
                ]
            },
        ]
    },
    tableData: [],
    pageData: {
        total: 0,
        pageSize: 10,
        pageNo: 1
    },
    query: () => {
        props.queryFunc(props.queryObj, base.pageData).then((res) => {
            if (res.status == 1) {
                base.tableData = res.data.data
                base.pageData.total = res.data.total
                base.pageData.pageSize = res.data.pageSize
                base.pageData.pageNo = res.data.pageNo
            } else {
                message("查询失败", "error")
            }
        })
    },
    handleSizeChange: (e) => {
        base.pageData.pageSize = e
        base.pageData.pageNo = 1
        base.query()
    },
    handleCurrentChange: (e) => {
        base.pageData.pageNo = e
        base.query()
    },
    select: (index, row) => {
        if (!(props.arrayObj.some(item => item.id == row.id))) {
            props.arrayObj.push(row)
        }
    },
    remove: (index, row) => {
        props.arrayObj.forEach((item, index, array) => {
            if (item.id == row.id) {
                array.splice(index, 1)
            }
        })
    }
})
onBeforeMount(() => {
    base.column.headers.unshift.apply(base.column.headers,props.headers)
    base.query()
})

defineExpose({
    base
})
</script>