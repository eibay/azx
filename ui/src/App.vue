<template>
  <div>
    <button @click="login">Login with Azure</button>
    <p v-if="userName">Logged in as: {{ userName }}</p>
  </div>
</template>

<script>
export default {
  data() {
    return {
      userName: ''
    }
  },
  methods: {
    async login() {
      try {
        const response = await fetch('http://localhost:8080/api/login');
        const data = await response.json();
        this.userName = data.userName;
      } catch (error) {
        console.error('Error logging in:', error);
      }
    }
  }
}
</script>


<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>
