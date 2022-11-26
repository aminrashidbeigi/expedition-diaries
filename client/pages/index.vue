<template>
   <div class="relative items-top justify-center min-h-screen bg-gray-100 sm:items-center sm:pt-0">
    <link href="https://cdn.jsdelivr.net/npm/tailwindcss@2.1.2/dist/tailwind.min.css" rel="stylesheet">
    <div class="max-w-4xl mx-auto sm:px-6 lg:px-8 container">
      <Header/>
      <div class="mt-2 bg-white overflow-hidden shadow sm:rounded-lg p-6">
        <h1 class="text-2xl leading-7 font-bold">
          Recently added expeditions
        </h1>
        <div class="mt-4 pt-4 text-gray-800 border-t border-dashed">
          <div v-if="!travels">
            No expedition found.
            <br>
            Help us by 
            <a 
              href="/add-travel"
              class="underline text-blue-600 hover:text-blue-800 visited:text-purple-600">
                suggesting an expedition
            </a> :)
          </div>
          <div v-for="travel in travels">
            <a :href="`/travels/`+travel.Slug" class="underline text-blue-600 hover:text-blue-800 visited:text-purple-600">
              <h2 class="font-bold text-xl">{{travel.Title}}</h2>
            </a>
            <br>

            <h3 class="font-bold">📅 Duration</h3>
            <div>
              <p>From {{travel.StartedAt? travel.StartedAt : "?"}} to {{travel.EndedAt? travel.EndedAt : "?"}}</p>
            </div>
            <br>
            
            <h3 class="font-bold">🤠 Travelers</h3>
            <div v-for="traveler in travel.Travelers">
              <p v-if="traveler.Link">
                - <a :href="traveler.Link" target="_blank" class="underline text-blue-600 hover:text-blue-800 visited:text-purple-600">
                  {{traveler.Name}}
                </a>
              </p>
              <p v-else>
                - {{traveler.Name}}
              </p>
            </div>
            <br>
        
            <div v-if="travel.Route">
              <h3 class="font-bold">📍 Route</h3>
              <div class="relative mx-auto w-80">
                <a :href="$config.baseImageURL+travel.Route" target="_blank">
                  <img :src="$config.baseImageURL+travel.Route" :alt="travel.Title + ' expedition route'" class="relative w-full h-full rounded-md shadow-md">
                </a>
              </div>
              <br>
            </div>
         
            <h3 class="font-bold">🌎 Countries of this travel</h3>
            <ul>
              <li v-for="country in travel.Countries">
                - <a :href="`/countries/` + country.Code" class="underline text-blue-600 hover:text-blue-800 visited:text-purple-600">
                  {{country.Name}}
                </a>
              </li> 
            </ul>
            <br>

            <h3 class="font-bold">📖 Resources</h3>
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
import Header from '../components/Header.vue'
import Footer from '../components/Footer.vue'

export default {
  data() {
    return {
      travels: {},
      title: "",
    }
  },
  head() {
    return {
      meta: [
        {
          hid: 'description',
          name: 'description',
          content: 'The journeys of explorers who have passed through lands throughout history'
        }
      ]
    }
  },
  components: {
    Header,
    Footer,
  },
  async asyncData({ $axios }) {
    const travels = await $axios.$get(process.env.baseAPI + '/travels')
    return { travels }
  },
}
</script>