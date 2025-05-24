<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="主键编码:" prop="id">
          <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="ConfigName:" prop="configName">
          <el-input v-model="formData.configName" :clearable="true"  placeholder="请输入ConfigName" />
       </el-form-item>
        <el-form-item label="ConfigKey:" prop="configKey">
          <el-input v-model="formData.configKey" :clearable="true"  placeholder="请输入ConfigKey" />
       </el-form-item>
        <el-form-item label="ConfigValue:" prop="configValue">
          <el-input v-model="formData.configValue" :clearable="true"  placeholder="请输入ConfigValue" />
       </el-form-item>
        <el-form-item label="ConfigType:" prop="configType">
          <el-input v-model="formData.configType" :clearable="true"  placeholder="请输入ConfigType" />
       </el-form-item>
        <el-form-item label="是否前台:" prop="isFrontend">
          <el-input v-model="formData.isFrontend" :clearable="true"  placeholder="请输入是否前台" />
       </el-form-item>
        <el-form-item label="Remark:" prop="remark">
          <el-input v-model="formData.remark" :clearable="true"  placeholder="请输入Remark" />
       </el-form-item>
        <el-form-item label="创建者:" prop="createBy">
          <el-input v-model.number="formData.createBy" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="更新者:" prop="updateBy">
          <el-input v-model.number="formData.updateBy" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="创建时间:" prop="createdAt">
          <el-date-picker v-model="formData.createdAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="最后更新时间:" prop="updatedAt">
          <el-date-picker v-model="formData.updatedAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="删除时间:" prop="deletedAt">
          <el-date-picker v-model="formData.deletedAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createSysConfig,
  updateSysConfig,
  findSysConfig
} from '@/api/demo/sysConfig'

defineOptions({
    name: 'SysConfigForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            id: undefined,
            configName: '',
            configKey: '',
            configValue: '',
            configType: '',
            isFrontend: '',
            remark: '',
            createBy: undefined,
            updateBy: undefined,
            createdAt: new Date(),
            updatedAt: new Date(),
            deletedAt: new Date(),
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findSysConfig({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createSysConfig(formData.value)
               break
             case 'update':
               res = await updateSysConfig(formData.value)
               break
             default:
               res = await createSysConfig(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
