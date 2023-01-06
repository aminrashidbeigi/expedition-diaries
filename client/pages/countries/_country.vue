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
          <div v-for="travel in countryTravels.Travels">
            <a :href="`/travels/`+travel.Slug" class="underline text-blue-600 hover:text-blue-800 visited:text-purple-600">
              <h2 class="font-bold text-xl">{{travel.Title}}</h2>
            </a>
            <div v-if="travel.Description" class="whitespace-pre-line">
              <p>
                {{ travel.Description }}
              </p>
            </div>  
            <br>

            <h2 class="font-bold">📅 Duration</h2>
            <div>
              <p>From {{travel.StartedAt? travel.StartedAt : "?"}} A.D. to {{travel.EndedAt? travel.EndedAt : "?"}} A.D.</p>
            </div>
            <br>
            
          
            <h2 class="font-bold">🤠 Explorers</h2>
            <div v-for="traveler in travel.Travelers">
              <p v-if="traveler.Link">
                - <a :href="traveler.Link" target="_blank" class="underline text-blue-600 hover:text-blue-800 visited:text-purple-600">
                  {{traveler.Name}}
                </a>
              </p>
              <p v-else>
                - {{traveler.Name}}
              </p>
              <div v-if="traveler.Image" class="relative mx-auto w-80">
                <img :src="$config.baseImageURL+traveler.Image" :alt="traveler.Name + ' image'" class="relative w-full h-full rounded-md shadow-md">
              </div>
            </div>
            <br>
        
            <div v-if="travel.Route">
              <h2 class="font-bold">📍 Route</h2>
              <div class="relative mx-auto w-80">
                <a :href="$config.baseImageURL+travel.Route" target="_blank">
                  <img :src="$config.baseImageURL+travel.Route" :alt="travel.Title + ' expedition route'" class="relative w-full h-full rounded-md shadow-md">
                </a>
              </div>
              <br>
            </div>
 
            <div class="flex space-x-2 justify-end">
              <a :href="`/travels/`+travel.Slug" target="_blank">
                <button type="button" class="inline-block px-6 py-2.5 bg-blue-500 text-white font-medium text-xs leading-tight uppercase rounded shadow-md hover:bg-blue-700 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:shadow-lg transition duration-150 ease-in-out">Read More</button>
              </a>
            </div>

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
