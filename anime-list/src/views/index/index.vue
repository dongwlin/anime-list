<script setup>
import AnimeCard from '@/views/index/components/AnimeCard.vue';
import useAnimeList from "@/store/modules/animeList.js";

useAnimeList().handleList();
</script>

<template>
  <div class="list-container">
    <template v-for="item in useAnimeList().data" :key="item.id">
      <template v-if="item.status">
        <AnimeCard
          :status="item.status"
          :img-src="item.img"
          :name="item.name"
          :type="item.type"
          :url="item.url"
          :dir="item.dir"
          :day="item.day"
          :description="item.description"
        ></AnimeCard>
      </template>
    </template>
  </div>
</template>

<style scoped>
@media screen and (min-width: 1360px) {
  .list-container {
  display: grid;
    grid-template-columns: repeat(4, 25%);
    justify-items: center;
    grid-row-gap: 1vh;
  }
}

@media screen and (min-width: 992px) and (max-width: 1360px) {
  .list-container {
    display: grid;
    grid-template-columns: repeat(3, auto);
    justify-items: center;
    grid-row-gap: 1vh;
  }
}

@media screen and (min-width: 768px) and (max-width: 992px) {
  .list-container {
    display: grid;
    grid-template-columns: repeat(2, 50%);
    justify-items: center;
    grid-row-gap: 1vh;
  }
}

@media screen and (max-width: 768px) {
  .list-container {
    display: grid;
    grid-template-columns: repeat(1, auto);
    justify-items: center;
    grid-row-gap: 1vh;
  }
}
</style>
