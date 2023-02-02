<template>
   <div class="relative items-top justify-center min-h-screen bg-gray-100 sm:items-center sm:pt-0">
    <link href="https://cdn.jsdelivr.net/npm/tailwindcss@2.1.2/dist/tailwind.min.css" rel="stylesheet">
    <div class="max-w-4xl mx-auto sm:px-6 lg:px-8 container">
      <Header/>
      <div class="mt-2 bg-white overflow-hidden shadow sm:rounded-lg p-6">
        <h1 class="text-2xl leading-7 font-bold">
          {{ countryTravels.Country.Name }} expeditions
        </h1>
        <div class="mt-4 pt-4 text-gray-800 border-t border-dashed">
          <div v-if="!countryTravels.Travels">
            No expedition has been submitted for {{ countryTravels.Country.Name }} yet.
            <br>
            Do you want to submit one? Click 
            <a 
              href="/add-travel"
              class="underline text-blue-600 hover:text-blue-800 visited:text-purple-600">
                here</a>.
          </div>
          <div v-for="(travel, index) in countryTravels.Travels">

            <ExpeditionPreview
              :title="travel.Title"
              :description="travel.Description"
              :route="travel.Route"
              :startedAt="travel.StartedAt"
              :endedAt="travel.EndedAt"
              :explorers="travel.Travelers"
              :slug="travel.Slug"
            ></ExpeditionPreview>

            <div v-if="index != countryTravels.Travels.length - 1" class="mt-4 pt-4 text-gray-800 border-t border-dashed"/>
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
      countryTravels: {
        Country: "",
        Travels: {}
      },
      travels: {},
      title: "",
    }
  },
  head() {
    return {
      title: this.countryTravels.Country.Name + ' expeditions',
      meta: [
        {
          hid: 'description',
          name: 'description',
          content: 'The journeys of explorers who have passed through ' + this.countryTravels.Country.Name + ' throughout history'
        }
      ]
    }
  },
  components: {
    Header,
    Footer,
  },
  async asyncData({ params, $axios }) {
    const countryTravels = await $axios.$get(process.env.baseAPI + '/country-travels/' + params.country)
    return { countryTravels }
  },
}
</script>
