import useAnimeList from "@/store/modules/animeList.js";
import useServer from '@/store/modules/server.js';

const useStore = () => ({
    useAnimeList,
    useServer
})

export default useStore;

