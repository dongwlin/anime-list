<template>
  <div class="listBox">
    <el-card
      v-for="item in animeList"
      :key="item.id"
      class="listItem"
    >
      <div class="itemId">id: {{item.id}}</div>
      <el-form label-width="60px">
        <el-form-item label="name">
          <el-input v-model="item.name"></el-input>
        </el-form-item>
        <el-form-item label="status">
          <el-select v-model="item.status">
            <el-option
              v-for="option in options.status"
              :key="option.value"
              :label="option.label"
              :value="option.value"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="type">
          <el-select v-model="item.type">
            <el-option
              v-for="option in options.type"
              :key="option.value"
              :label="option.label"
              :value="option.value"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="dir">
          <el-input v-model="item.dir"></el-input>
        </el-form-item>
        <el-form-item label="url">
          <el-input v-model="item.url"></el-input>
        </el-form-item>
        <el-form-item label="img">
          <el-input v-model="item.img" disabled></el-input>
        </el-form-item>
        <el-form-item label="day">
          <el-select v-model="item.day">
            <el-option
              v-for="option in options.day"
              :key="option.value"
              :label="option.label"
              :value="option.value"
            ></el-option>
          </el-select>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script setup>
import {onMounted, reactive, ref} from "vue";
import {useAnimeOptions} from "@/store/index.js";
import request from "@/utils/axiosInstance.js";

const animeList = reactive([]);
const loading = ref(true);

const options = useAnimeOptions();

const getAnimeList = () => {
  request.get(
      '/animeList',
      {
        params: {

        }
      }
  ).then(response => {
    animeList.push(...response.data);
  }).catch(error => {
    console.log(error);
  });
  loading.value = false;
}

onMounted(() => {
  getAnimeList();
})
</script>

<style scoped>
.listBox {
  display: flex;
  justify-content: space-between;
  flex-wrap: wrap;
}

.listItem {
  margin-bottom: 20px;
}

.itemId {
  color: #999;
  font-size: 12px;
  margin-bottom: 5px;
}
</style>
