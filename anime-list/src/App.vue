<script setup>
import TitleBar from "@/components/TitleBar.vue"
</script>

<template>
  <TitleBar></TitleBar>
  <router-view v-slot="{Component, route}">
<!--    <component -->
<!--        :is="Component" -->
<!--        :key="route.path"-->
<!--    >-->
<!--    </component>-->
    <component
        :is="Component"
    >
    </component>
  </router-view>
</template>

<style scoped>

</style>
