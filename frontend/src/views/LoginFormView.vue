<template>
  <div class="login-form">
    <div>
      <h1 class="title">ログイン</h1>
      <hr class="title-divider" />
    </div>
    <div class="login-form">
      <input
        v-model="state.email"
        type="text"
        class="login-input-text"
        required
        placeholder="test@example.com"
      />
      <input
        v-model="state.password"
        type="password"
        class="login-input-text"
        required
      />
      <button @click="postLogin" class="login-button" type="button">
        送信
      </button>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive } from "vue";
import { useRouter } from "vue-router";
import axios from "axios";

const baseURL = "http://localhost:8080/";

export default defineComponent({
  name: "LoginForm",
  setup() {
    const state = reactive({
      email: "",
      password: "",
    });
    const router = useRouter();

    const postLogin = async () => {
      await axios.post(`${baseURL}api/v1/login`, {
        email: state.email,
        password: state.password,
      });
      router.push("/");
    };

    return {
      state,
      postLogin,
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
.login-form {
  margin: 100px;
}
.login-input-text {
  margin: 10px;
  width: 200px;
  height: 50px;
  font-size: 150%;
  border: 2px solid #bbb;
}
.login-button {
  margin: 10px;
  width: 100px;
  height: 56px;
  font-size: 150%;
  border: 2px solid #bbb;
  background: transparent;
  color: #444;
}
</style>
