<template>
  <div class="w-full flex flex-col gap-4 bg-zinc-800 p-8">
    <h1 class="font-bold text-3xl">Mamuro Email</h1>
    <div class="flex flex-row gap-2 h-12 rounded-md">
      <input type="text" v-model="search" @keyup.enter="searchMails" class="w-full h-full bg-zinc-700 rounded-md text-white px-6" placeholder="Search">
      <button @click="searchMails" class="h-full w-12 flex justify-center ">
        <span v-if="loading" class="material-symbols-outlined self-center animate-spin">loop</span>
        <span v-else class="material-symbols-outlined self-center">search</span>
      </button>
    </div>
    <div class="text-md flex">
    </div>
    <h1 v-if="error" class="text-xl text-red-600">Emails not found</h1>
    <h1 v-if="missing" class="text-xl text-red-600"> Missing arguments (Search)</h1>
  </div>
</template>

<script>
  import store from '../../store/mails';

  export default {
    name: 'Bar',
    data() {
      return {
        search: '',
        missing: false
      };
    },
    computed: {
      error() {
        return store.state.error;
      },
      loading() {
        return store.state.loadingMSG;
      }
    },
    methods: {
      searchMails() {
        if (this.search.trim() !== '') {
          store.dispatch('changeName', this.name);
          store.dispatch('searchMails', this.search);
        } else {
          this.missing = true;
          setTimeout(() => {
            this.missing = false;
          }, 3000);
        }
      },
    },
  };
</script>