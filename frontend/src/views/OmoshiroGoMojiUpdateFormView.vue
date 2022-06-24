<template>
  <div class="omoshiro-go-moji-update-form">
    <div>
      <h1 class="title">5文字</h1>
      <hr class="title-divider" />
    </div>
    <div class="omoshiro-go-moji-form">
      <input
        v-model="state.name"
        type="text"
        class="omoshiro-go-moji-input-name"
        required
        minlength="5"
        maxlength="5"
        placeholder="とちくるう"
      />
      <button
        @click="updateOmoshiroGoMoji"
        class="omoshiro-go-moji-button"
        type="button"
      >
        送信
      </button>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive } from "vue";
import { useRouter, useRoute } from "vue-router";
import axios from "axios";

const baseURL = "http://localhost:3000/";

interface OmoshiroGoMoji {
  name: string;
  id: number;
}

export default defineComponent({
  name: "OmoshiroGoMojiUpdateForm",
  setup() {
    const state = reactive({
      name: "",
      id: undefined,
    });
    const router = useRouter();
    const route = useRoute();

    const getOmoshiroGoMoji = async () => {
      const res = await axios.get(
        `${baseURL}api/v1/omoshiro_go_moji/${route.params.id}`
      );
      const omoshiroGoMoji: OmoshiroGoMoji = res.data.data;
      state.name = omoshiroGoMoji.name;
      state.id = omoshiroGoMoji.id;
    };

    getOmoshiroGoMoji();

    const updateOmoshiroGoMoji = async () => {
      await axios.patch(`${baseURL}api/v1/omoshiro_go_moji/${state.id}`, {
        name: state.name,
      });
      router.push("/");
    };

    return {
      state,
      updateOmoshiroGoMoji,
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
.omoshiro-go-moji-form {
  margin: 100px;
}
.omoshiro-go-moji-input-name {
  margin: 10px;
  width: 200px;
  height: 50px;
  font-size: 150%;
  border: 2px solid #bbb;
}
.omoshiro-go-moji-button {
  margin: 10px;
  width: 100px;
  height: 56px;
  font-size: 150%;
  border: 2px solid #bbb;
  background: transparent;
  color: #444;
}
</style>
