<script setup>
import { config } from "@/views/setting/edit/config.js";
import {toType, toDay} from "@/utils/toValue.js";
import useAnimeList from '@/store/modules/animeList.js';
import { disable } from "@/api/disable.js";
import { enable } from "@/api/enable.js";
import {ElMessage} from "element-plus";
import {defineAsyncComponent, ref} from "vue";
import {deleteAnime} from "@/api/delete.js";

const EditUpdateDialog = defineAsyncComponent(() => import('@/views/setting/edit/components/EditUpdateDialog.vue'));
const list = useAnimeList();
const updateDialogVisible = ref(false);
const readId = ref('');

const handleUpdateDialog = (id) => {
  readId.value = id;
  updateDialogVisible.value = true;
}

const closeUpdateDialog = () => {
  updateDialogVisible.value = false;
}

const handleDelete = async (id) => {
  await deleteAnime(id)
      .then(res => {
        if (res.code === 200) {
          list.handleList();
          ElMessage({
            message: "success",
            type: "success"
          });
        }
        else {
          ElMessage({
            message: "failed",
            type: "error"
          });
        }
      })
      .catch(error => {
        ElMessage({
          message: "failed",
          type: "error"
        });
      });
}

const requestDisable = async (id) => {
  await disable(id)
      .then(res => {
        if (res.code === 200) {
          ElMessage({
            message: "success",
            type: "success"
          });
        }
        else {
          ElMessage({
            message: "failed",
            type: "error"
          });
        }
      })
      .catch(error => {
        ElMessage({
          message: "failed",
          type: "error"
        });
      });
}

const requestEnable = async (id) => {
  await enable(id)
      .then(res => {
        if (res.code === 200) {
          ElMessage({
            message: "success",
            type: "success"
          });
        }
        else {
          ElMessage({
            message: "failed",
            type: "error"
          });
        }
      })
      .catch(error => {
        ElMessage({
          message: "failed",
          type: "error"
        });
      });
}

const statusSwitch = (id, value) => {
  if (value)
  {
    requestEnable(id)
    console.log(id,'true');
  }
  else {
    requestDisable(id);
    console.log(id,'false');
  }
}

list.handleList();
</script>

<template>
  <el-card>
    <el-table :data="list.data" height="65vh">
      <el-table-column
          v-for="(item, index) in config.tableHeader"
          :key="index"
          :prop="item.prop"
          :label="item.label"
          :width="item.width"
          :min-width="item.minWidth"
      >
        <template v-slot="{row}" v-if="item.prop === 'name'">
          <el-text truncated>{{ row.name }}</el-text>
        </template>
        <template v-slot="{row}" v-else-if="item.prop === 'type'">
          {{ toType(row.type) }}
        </template>
        <template v-slot="{row}" v-else-if="item.prop === 'day'">
          {{ toDay(row.day )}}
        </template>
        <template v-slot="{row}" v-else-if="item.prop === 'status'">
          <el-switch v-model="row.status" @change="statusSwitch(row.id, row.status)"></el-switch>
        </template>
        <template v-slot="{row}" v-else-if="item.prop === 'operations'">
          <el-button size="small" @click="handleUpdateDialog(row.id)">Edit</el-button>
          <el-button size="small" type="danger" @click="handleDelete(row.id)">Delete</el-button>
        </template>
      </el-table-column>
    </el-table>
  </el-card>

  <EditUpdateDialog :visible="updateDialogVisible" :id="readId" @close="closeUpdateDialog"></EditUpdateDialog>
</template>

<style scoped>

</style>
