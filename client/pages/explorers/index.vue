<template>
   <div class="relative items-top justify-center min-h-screen bg-gray-100 sm:items-center sm:pt-0">
    <link href="https://cdn.jsdelivr.net/npm/tailwindcss@2.1.2/dist/tailwind.min.css" rel="stylesheet">
    <div class="max-w-4xl mx-auto sm:px-6 lg:px-8 container">
      <Header/>
      <div class="mt-2 bg-white overflow-hidden shadow sm:rounded-lg p-6">
        <h1 class="text-2xl leading-7 font-bold">
          Explorers
        </h1>
        <div class="mt-4 pt-4 text-gray-800 border-t border-dashed">
          
          <ul v-for="explorer in explorers">
            <li class="font-bold">
              - <a :href="`/explorers/` + explorer.Slug.String" class="underline text-blue-600 hover:text-blue-800 visited:text-purple-600" target="_blank">
                {{ explorer.Name }}
                </a>
            </li>
          </ul>
        
        </div>
      </div>
      <Footer/>
    </div>
  </div>
</template>

<script>
import Header from '../../components/Header.vue'
import Footer from '../../components/Footer.vue'

export default {
  data() {
    return {
      explorers: {},
    }
  },
  head() {
    return {
      title: 'Expedition Diaries - Explorers',
      meta: [
        {
          hid: 'description',
          name: 'description',
          content: 'Explorers who have passed through lands throughout history'
        }
      ]
    }
  },
  components: {
    Header,
    Footer,
  },
  async asyncData({ $axios }) {
    const explorers = await $axios.$get(process.env.baseAPI + '/travelers')
    return { explorers }
  },
}
</script>