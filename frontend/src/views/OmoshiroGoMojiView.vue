<template>
  <div class="omoshiro-go-moji">
    <header>
      <h1 class="title">おもしろ5文字</h1>
      <hr class="title-divider" />
    </header>
    <main class="cards">
      <div class="card">
        <div class="card-content">
          <p class="card-text">どんとこい</p>
          <router-link to="#">編集</router-link> |
          <router-link to="#">削除</router-link>
        </div>
        <hr class="card-divider" />
      </div>
      <div class="card">
        <div class="card-content">
          <p class="card-text">どんとこい</p>
          <router-link to="#">編集</router-link> |
          <router-link to="#">削除</router-link>
        </div>
        <hr class="card-divider" />
      </div>
      <div class="card">
        <div class="card-content">
          <p class="card-text">どんとこい</p>
          <router-link to="#">編集</router-link> |
          <router-link to="#">削除</router-link>
        </div>
        <hr class="card-divider" />
      </div>
      <div class="card">
        <div class="card-content">
          <p class="card-text">どんとこい</p>
          <router-link to="#">編集</router-link> |
          <router-link to="#">削除</router-link>
        </div>
        <hr class="card-divider" />
      </div>
    </main>
  </div>
</template>

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
