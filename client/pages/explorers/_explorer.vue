<template>
   <div class="relative items-top justify-center min-h-screen bg-gray-100 sm:items-center sm:pt-0">
    <link href="https://cdn.jsdelivr.net/npm/tailwindcss@2.1.2/dist/tailwind.min.css" rel="stylesheet">
    <div class="max-w-4xl mx-auto sm:px-6 lg:px-8 container">
      <Header/>
      <div class="mt-2 bg-white overflow-hidden shadow sm:rounded-lg p-6">
        <h1 class="text-2xl leading-7 font-bold">
          {{ traveler.Name }} expeditions
        </h1>
        <div class="mt-4 pt-4 text-gray-800 border-t border-dashed">
          <div v-for="(travel, index) in travels">

            <ExpeditionPreview
              :title="travel.Title"
              :description="travel.Description"
              :route="travel.Route"
              :startedAt="travel.StartedAt"
              :endedAt="travel.EndedAt"
              :explorers="travel.Travelers"
              :slug="travel.Slug"
            ></ExpeditionPreview>

            <div v-if="index != travels.length - 1" class="mt-4 pt-4 text-gray-800 border-t border-dashed"/>
          </div>
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
      travels: {},
      traveler: "",
    }
  },
  head() {
    return {
      title: this.traveler.Name + ' Expeditions',
      meta: [
        {
          hid: 'description',
          name: 'description',
          content: 'The expeditions of' + this.traveler.Name
        }
      ]
    }
  },
  components: {
    Header,
    Footer,
  },
  async asyncData({ params, $axios }) {
    const traveler = await $axios.$get(process.env.baseAPI + '/travelers/' + params.explorer)
    const travels = await $axios.$get(process.env.baseAPI + '/travelers/' + params.explorer + '/travels')

    return { travels, traveler }
  },
}
</script>
