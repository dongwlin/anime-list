<template>
  <transition name="fade" mode="out-in" appear>
    <el-container>
      <el-aside>
        <el-menu>
          <el-menu-item index="1" @click="ToSettingServer">
            <span>Server</span>
          </el-menu-item>
          <el-menu-item index="2" @click="ToSettingEdit">
            <span>Edit</span>
          </el-menu-item>
          <el-menu-item index="3">
            <span>About</span>
          </el-menu-item>
        </el-menu>
      </el-aside>
      <el-main>
        <router-view></router-view>
      </el-main>
    </el-container>
  </transition>
</template>

<script setup>
import { useRouter } from "vue-router"

const router = useRouter();

const ToSettingServer = () => {
  router.push('/setting/server');
}

const ToSettingEdit = () => {
  router.push('/setting/edit');
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
</style>
