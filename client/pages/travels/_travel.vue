<template>
  <div class="relative items-top justify-center min-h-screen bg-gray-100 sm:items-center sm:pt-0">
    <link href="https://cdn.jsdelivr.net/npm/tailwindcss@2.1.2/dist/tailwind.min.css" rel="stylesheet">
    <div class="max-w-4xl mx-auto sm:px-6 lg:px-8 container">
      <Header/>

      <Expedition
        :title="travel.Title"
        :description="travel.Description"
        :route="travel.Route"
        :startedAt="travel.StartedAt"
        :endedAt="travel.EndedAt"
        :explorers="travel.Travelers"
        :resources="travel.Resources"
        :countries="travel.Countries"
      ></Expedition>

      <Footer/>
    </div>
  </div>
</template>

<script>
import Header from '../../components/Header.vue'
import Footer from '../../components/Footer.vue'
import Expedition from '../../components/Expedition.vue'

export default {
  data() {
    return {
      travel: {},
      title: "",
    }
  },
  head() {
    return {
      title: this.travel.Title,
      meta: [
        {
          hid: 'description',
          name: 'description',
          content: this.travel.Description
        }
      ]
    }
  },
  components: {
    Header,
    Footer,
    Expedition,
  },
  
  async asyncData({ params, error, $axios }) {
    let travel = null;
    try{
      travel = await $axios.$get(process.env.baseAPI + '/travels/' + params.travel)
    } catch (err) {
      if (err.response){
        error({statusCode: err.response.status, message: err.response.data})
      }
    }
    return { travel }
  },
}
</script>
