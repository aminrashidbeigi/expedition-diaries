<template>
   <div class="relative items-top justify-center min-h-screen bg-gray-100 sm:items-center sm:pt-0">
    <link href="https://cdn.jsdelivr.net/npm/tailwindcss@2.1.2/dist/tailwind.min.css" rel="stylesheet">
    <div class="max-w-4xl mx-auto sm:px-6 lg:px-8 container">
      <Header/>
      <div class="mt-2 bg-white overflow-hidden shadow sm:rounded-lg p-6">
        <h1 class="text-2xl leading-7 font-bold">
          {{ travel.Title }}
        </h1>
        <div class="mt-4 pt-4 text-gray-800 border-t border-dashed">
            <div v-if="travel.Description" class="whitespace-pre-line"><p>{{ travel.Description }}</p></div> 
            <br>

            <h2 class="font-bold">📅 Duration</h2>
            <div>
              <p>From {{travel.StartedAt? travel.StartedAt : "?"}} to {{travel.EndedAt? travel.EndedAt : "?"}}</p>
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
