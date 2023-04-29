<template>
  <transition name="fade" mode="out-in" appear>
    <el-container>
      <el-aside>
        <el-menu
          :router="true"
          :default-active="menuDefaultActive"
        >
          <el-menu-item index="/setting/server">
            <span class="menuTitle">Server</span>
          </el-menu-item>
          <el-menu-item index="/setting/edit">
            <span class="menuTitle">Edit</span>
          </el-menu-item>
          <el-menu-item index="/setting/about">
            <span class="menuTitle">About</span>
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
const menuDefaultActive = location.hash.substring(1, location.hash.length);
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

.menuTitle {
  padding-left: 20px;
}
</style>
