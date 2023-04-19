<template>
  <transition name="fade" mode="out-in" appear>
    <div class="animeListContainerBase">
      <div class="animeListContainer">
        <template v-for="item in animeListA" :key="item.id">
          <el-row v-if="item.status">
            <el-col :span="6"></el-col>
            <el-col :span="16" v-if="item.type === 0">
              <el-card shadow="hover" @click="openUrl(item.url)">{{ item.name }}</el-card>
            </el-col>
            <el-col :span="16" v-else-if="item.type === 1">
              <el-card shadow="hover" @click="openDir(item.dir)">{{ item.name }}</el-card>
            </el-col>
            <el-col :span="16" v-else>
              <el-card shadow="hover">{{ item.name }}</el-card>
            </el-col>
            <el-col :span="2"></el-col>
          </el-row>
        </template>
      </div>
      <div class="animeListContainer">
        <template v-for="item in animeListB" :key="item.id">
          <el-row v-if="item.status">
            <el-col :span="2"></el-col>
            <el-col :span="16" v-if="item.type === 0">
              <el-card shadow="hover" @click="openUrl(item.url)">{{ item.name }}</el-card>
            </el-col>
            <el-col :span="16" v-else-if="item.type === 1">
              <el-card shadow="hover" @click="openDir(item.dir)">{{ item.name }}</el-card>
            </el-col>
            <el-col :span="16" v-else>
              <el-card shadow="hover">{{ item.name }}</el-card>
            </el-col>
            <el-col :span="6"></el-col>
          </el-row>
        </template>
      </div>
    </div>
  </transition>
</template>

<script setup>
import request from "@/utils/axiosInstance.js";
import {ElNotification} from 'element-plus'
import {onMounted, reactive} from "vue";

let animeListA = reactive([]);
let animeListB = reactive([]);

const openUrl = (url) => {
  open(url, '_blank');
}

const openDir = async (folderName) => {
  request.get(
      '/open',
      {
        params: {
          folder: folderName
        }
      }
  ).then(response => {
    if (response.code === 200)
    {
      ElNotification({
        title: 'Success',
        message: 'Open folder success.',
        type: 'success',
        position: 'bottom-right'
      });
    }
    else if (response.code === 404)
    {
      ElNotification({
        title: 'Error',
        message: 'Open folder/url fail.',
        type: 'error',
        position: 'bottom-right'
      });
    }
  }).catch(error => {
    console.log(error);
  });
}

const getAnimeListA = async () => {
  request.get(
      '/animeListA',
      {
        params: {

        }
      }
  ).then(response => {
    animeListA.push(...response.data);
  }).catch(error => {
    console.log(error);
  });
}

const getAnimeListB = async () => {
  request.get(
      '/animeListB',
      {
        params: {

        }
      }
  ).then(response => {
    animeListB.push(...response.data);
  }).catch(error => {
    console.log(error);
  });
}

onMounted(() => {
  getAnimeListA();
  getAnimeListB();
});
</script>

<style scoped>
.el-row:first-child {
  margin-top: 20px;
}

.el-row {
  margin-bottom: 20px;
}

.el-row:last-child {
  margin-bottom: 0;
}

.fade-enter-active,
.fade-leave-active {
  transition: all 200ms ease-out;
}

.fade-enter-from {
  opacity: 0;
  transform: translateX(30px);
}

.fade-leave-to {
  opacity: 0;
  transform: translateX(-30px);
}

.animeListContainerBase {
  display: flex;
  justify-content: space-between;
}

.animeListContainer {
  width: 50vw;
}
</style>
