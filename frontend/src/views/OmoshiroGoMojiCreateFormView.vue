<template>
  <div class="omoshiro-go-moji-create-form">
    <div>
      <h1 class="title">おもしろ5文字</h1>
      <hr class="title-divider" />
    </div>
    <div class="omoshiro-go-moji-form">
      <input
        type="text"
        name="name"
        class="omoshiro-go-moji-input-name"
        required
        minlength="5"
        maxlength="5"
        placeholder="とちくるう"
      />
      <button class="omoshiro-go-moji-button" type="button">送信</button>
    </div>
  </div>
</template>

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
