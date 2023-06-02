<script setup>
import { Search, DocumentAdd } from '@element-plus/icons-vue';
import {defineAsyncComponent, ref} from 'vue';

const CreateDialog = defineAsyncComponent(() => import('@/views/setting/edit/components/EditCreateDialog.vue'));
const keyword = ref('');
const CreateDialogVisible = ref(false);

const handleCreateDialog = () => {
  CreateDialogVisible.value = true;
}

const closeCreateDialog = () => {
  CreateDialogVisible.value = false;
}
</script>

<template>
  <el-card class="container">
    <el-row :gutter="20">
      <el-col :span="6" :offset="1">
        <el-input v-model="keyword" placeholder="Search" disabled></el-input>
      </el-col>
      <el-col :span="8">
        <el-button type="primary" :icon="Search" disabled>Search</el-button>
        <el-button type="primary" :icon="DocumentAdd" @click="handleCreateDialog">Create</el-button>
      </el-col>
    </el-row>
  </el-card>

  <CreateDialog :visible="CreateDialogVisible" @close="closeCreateDialog"></CreateDialog>
</template>

<style scoped>
.container {
  margin-bottom: 20px;
}
</style>
