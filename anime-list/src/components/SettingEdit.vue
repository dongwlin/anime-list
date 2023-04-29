<template>
  <transition name="fade" mode="out-in" appear>
    <div>
      <el-row>
        <el-col :span="3"></el-col>
        <el-col :span="18">
          <el-card>
            <div class="setting-item">
              <span>Add Anime</span>
              <el-button
                  :type="'primary'"
                  @click="dialogAdd = true"
              >Add</el-button>
              <el-dialog v-model="dialogAdd" title="Add Anime" @closed="clearAdd">
                <el-form label-width="60px">
                  <el-form-item label="name">
                    <el-input v-model="formAdd.name" class="inputBox"></el-input>
                  </el-form-item>
                  <el-form-item label="status">
                    <el-select v-model="formAdd.status" default-first-option>
                      <el-option
                          v-for="item in options.status"
                          :key="item.value"
                          :label="item.label"
                          :value="item.value"
                      ></el-option>
                    </el-select>
                  </el-form-item>
                  <el-form-item label="type">
                    <el-select v-model="formAdd.type" default-first-option>
                      <el-option
                          v-for="item in options.type"
                          :key="item.value"
                          :label="item.label"
                          :value="item.value"
                      ></el-option>
                    </el-select>
                  </el-form-item>
                  <el-form-item label="dir">
                    <el-input v-model="formAdd.dir" class="inputBox"  @change="replaceBackslash(formAdd)"></el-input>
                  </el-form-item>
                  <el-form-item label="url">
                    <el-input v-model="formAdd.url" class="inputBox"></el-input>
                  </el-form-item>
                  <el-form-item label="img">
                    <el-input class="inputBox" v-model="formAdd.img" disabled placeholder="none"></el-input>
                  </el-form-item>
                  <el-form-item label="day">
                    <el-select v-model="formAdd.day" default-first-option>
                      <el-option
                          v-for="item in options.day"
                          :key="item.value"
                          :label="item.label"
                          :value="item.value"
                      ></el-option>
                    </el-select>
                  </el-form-item>
                </el-form>
                <template #footer>
                  <el-button @click="cancelAdd">Cancel</el-button>
                  <el-button :type="'primary'" @click="addAnime">Confirm</el-button>
                </template>
              </el-dialog>
            </div>
          </el-card>
        </el-col>
        <el-col :span="3"></el-col>
      </el-row>
      <el-row>
        <el-col :span="3"></el-col>
        <el-col :span="18">
          <Suspense>
            <template #default>
              <SettingEditList></SettingEditList>
            </template>
            <template #fallback>
              <div class="loadingBox">loading</div>
            </template>
          </Suspense>
        </el-col>
        <el-col :span="3"></el-col>
      </el-row>
    </div>
  </transition>
</template>

<script setup>
import {reactive, ref} from "vue";
import request from "@/utils/axiosInstance.js";
import { useAnimeOptions } from '@/store/index.js';
import SettingEditList from '@/components/SettingEditList.vue';

const dialogAdd = ref(false);
const formAdd = reactive({
  name: '',
  status: true,
  type: -1,
  dir: '',
  url: '',
  img: 'none',
  day: -1
});

const options = useAnimeOptions();

const clearAdd = () => {
  formAdd.name = '';
  formAdd.status = true;
  formAdd.type = -1;
  formAdd.dir = '';
  formAdd.url = '';
  formAdd.img = 'none';
  formAdd.day = -1;
}

const cancelAdd = () => {
  clearAdd();
  dialogAdd.value = false;
}

const replaceBackslash = (obj) => {
  obj.dir = obj.dir.replaceAll('\\', '/');
}

const addAnime = () => {
  let formData = new FormData();
  if (formAdd.name.length !== 0)
  {
    formData.append('name', formAdd.name);
  }
  formData.append('status', formAdd.status);
  formData.append('type', formAdd.type);
  formData.append('dir', formAdd.dir);
  formData.append('url', formAdd.url);
  if (formAdd.img !== 'none')
  {
    formData.append('img', formAdd.img);
  }
  formData.append('day', formAdd.day);

  request.post(
      '/add',
      formData,
      {
        headers: {
          'Content-Type': 'multipart/form-data;charset=utf-8'
        }
      }
  ).then(response => {
    console.log(response);
  }).catch(error => {
    console.log(error);
  });

  clearAdd();
  dialogAdd.value = false;
}
</script>

<style scoped>
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

.el-row:first-child {
  margin-top: 20px;
}

.el-row {
  margin-bottom: 20px;
}

.el-row:last-child {
  margin-bottom: 0;
}

.setting-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.inputBox {
  width: 550px;
}
</style>
