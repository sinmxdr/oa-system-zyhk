<template>
    <div v-if="user().my.pids.includes('25')">
        <el-divider content-position="left">
            <h2>客户管理</h2>
        </el-divider>
        <customerManage ref="customerPage" />
    </div>
    <div style="margin-top: 30px;" v-if="user().my.pids.includes('27')">
        <el-divider content-position="left">
            <h2>客户公司管理</h2>
        </el-divider>
        <companyManage :handleAddCustomer="resetCustomerPage" />
    </div>
</template>

<script setup>
import { user } from '@/pinia/modules/user'
import { ref } from 'vue'

import customerManage from './CustomerManage.vue'
import companyManage from './CompanyManage.vue'

const customerPage = ref(null)

function resetCustomerPage() {
    customerPage.value.base.query()
}
</script>