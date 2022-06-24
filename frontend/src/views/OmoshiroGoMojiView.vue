<template>
  <div class="omoshiro-go-moji">
    <header>
      <h1 class="title">5文字</h1>
      <hr class="title-divider" />
    </header>
    <main class="cards">
      <div
        class="card"
        v-for="omoshiroGoMoji in state.omoshiroGoMojiList"
        :key="omoshiroGoMoji.uuid"
      >
        <div class="card-content">
          <p class="card-text">{{ omoshiroGoMoji.name }}</p>
          <router-link
            :to="{
              name: 'omoshiro_go_moji_update_form',
              params: { id: omoshiroGoMoji.id },
            }"
          >
            編集
          </router-link>
          <a @click="deleteOmoshiroGomoji(omoshiroGoMoji.id)">削除</a>
        </div>
        <hr class="card-divider" />
      </div>
    </main>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive } from "vue";
import axios from "axios";

const baseURL = "http://localhost:3000/";

interface OmoshiroGoMoji {
  name: string;
  id: number;
}

const defaultOmoshiroGoMojiList: OmoshiroGoMoji[] = [];

export default defineComponent({
  name: "OmoshiroGoMoji",
  setup() {
    const state = reactive({
      omoshiroGoMojiList: defaultOmoshiroGoMojiList,
    });

    const getOmoshiroGoMojiList = async () => {
      const res = await axios.get(`${baseURL}api/v1/omoshiro_go_moji/list`);
      const omoshiroGoMojiList: OmoshiroGoMoji[] = res.data.data;
      state.omoshiroGoMojiList = omoshiroGoMojiList;
    };

    getOmoshiroGoMojiList();

    const deleteOmoshiroGomoji = async (id: number) => {
      await axios.delete(`${baseURL}api/v1/omoshiro_go_moji/${id}`);
      const res = await axios.get(`${baseURL}api/v1/omoshiro_go_moji/list`);
      const omoshiroGoMojiList: OmoshiroGoMoji[] = res.data.data;
      state.omoshiroGoMojiList = omoshiroGoMojiList;
    };

    return {
      state,
      deleteOmoshiroGomoji,
    };
  },
});
</script>

<style lang="scss">
.title {
  font-size: 40px;
}
hr.title-divider {
  border-top: 2px solid #bbb;
}
hr.card-divider {
  border-top: 1px solid #bbb;
}
.cards {
  margin-top: 50px;
  margin-right: 100px;
  margin-left: 100px;
}
.card-content {
  display: flex;
  justify-content: center;
  align-items: center;
}
.card {
  margin-top: 100px;
}
.card-text {
  font-size: 32px;
  font-weight: bold;
  letter-spacing: 16px;
}
a {
  font-weight: bold;
  color: #2c3e50;
  margin: 5px;
}
</style>
