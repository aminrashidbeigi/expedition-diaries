<template>
   <div class="relative flex items-top justify-center min-h-screen bg-gray-100 sm:items-center sm:pt-0">
    <link href="https://cdn.jsdelivr.net/npm/tailwindcss@2.1.2/dist/tailwind.min.css" rel="stylesheet">
    <div class="max-w-4xl mx-auto sm:px-6 lg:px-8 container">
      <Header/>
      <div class="mt-8 bg-white overflow-hidden shadow sm:rounded-lg p-6">
        <h2 class="text-2xl leading-7 font-bold">
          {{ countryTravels.Country.Name }} expeditions
        </h2>
        <div class="mt-4 pt-4 text-gray-800 border-t border-dashed">
          <div v-if="!countryTravels.Travels">
            No expedition found for {{ countryTravels.Country.Name }}.
            <br>
            Help us by 
            <a 
              href="/add-travel"
              class="underline text-blue-600 hover:text-blue-800 visited:text-purple-600">
                suggesting an expedition
            </a> :)
          </div>
          <div v-for="travel in countryTravels.Travels">
            <h2 class="font-bold text-xl">{{travel.Title}}</h2>
            <br>

            <h2 class="font-bold">📅 Duration</h2>
            <div>
              <p>From {{travel.StartedAt? travel.StartedAt : "?"}} to {{travel.EndedAt? travel.EndedAt : "?"}}</p>
            </div>
            <br>

           
            <h2 class="font-bold">🤠 Travelers</h2>
            <div v-for="traveler in travel.Travelers">
              <p>
                - <a :href="traveler.Link" target="_blank" class="underline text-blue-600 hover:text-blue-800 visited:text-purple-600">
                  {{traveler.Name}}
                </a>
              </p>
            </div>
            <br>

            <h2 class="font-bold">🌎 Countries of this travel</h2>
            <ul>
              <li v-for="country in travel.Countries">
                - <a :href="`/countries/` + country.Code" class="underline text-blue-600 hover:text-blue-800 visited:text-purple-600">
                  {{country.Name}}
                </a>
              </li> 
            </ul>
            <br>

            <h2 class="font-bold">📖 Resources</h2>
            <div v-for="resource in travel.Resources">
              <p class="font-meduim">
                <a :href="resource.Link" target="_blank" class="underline text-blue-600 hover:text-blue-800 visited:text-purple-600">
                  {{resource.Title}}
                </a>
              </p>
              <div class="relative mx-auto w-60 h-80">
                <a :href="resource.Link" target="_blank">
                  <img :src="$config.baseImageURL+resource.Image" :alt="resource.Title" class="relative w-full h-full rounded-md shadow-md">
                </a>
              </div>
              <br>
            </div>

            
            <div class="mt-4 pt-4 text-gray-800 border-t border-dashed"/>
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
  data: {
    countryTravels: {
      Country: "",
      Travels: {}
    },
    travels: {}
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
