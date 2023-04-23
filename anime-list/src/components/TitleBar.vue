<template>
  <el-row>
    <el-col :span="24">
      <div class="title_bar">
        <div class="title">Anime List</div>
        <transition name="slide-up" mode="out-in">
          <el-button
              v-if="!isSetting"
              plain
              @click="ToSetting"
              class="btn"
          >
            <el-icon><Setting /></el-icon>
          </el-button>
          <el-button
              v-else
              plain
              @click="ToHome"
              class="btn"
          >
            <el-icon><Back /></el-icon>
          </el-button>
        </transition>
      </div>
    </el-col>
  </el-row>
</template>

<script setup>
import { useRouter } from "vue-router";
import { ref } from "vue";

const router = useRouter();

let isSetting = ref(false);
let location_hash = location.hash;
if (location_hash.length >= 9 && location_hash.substring(0, 9) === '#/setting')
{
  isSetting.value = true;
}

const ToSetting = () => {
  router.push('/setting');
  isSetting.value = true;
}

const ToHome = () => {
  router.push('/');
  isSetting.value = false;
}
</script>

<style scoped>
.title_bar {
  height: 42px;
  font-size: 22px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px 40px;
}

.title {
  font-size: 36px;
  font-style: italic;
}

.btn {
  width: 60px;
}

.slide-up-enter-active,
.slide-up-leave-active {
  transition: all 200ms ease-out;
}

.slide-up-enter-from {
  opacity: 0;
  transform: translateX(30px);
}

.slide-up-leave-to {
  opacity: 0;
  transform: translateX(-30px);
}
</style>
