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

          <!-- <section id="prev-next">
            <nuxt-link :to="prevLink">
              <button type="button" class="inline-block px-6 py-2.5 bg-blue-500 text-white font-medium text-xs leading-tight uppercase rounded shadow-md hover:bg-blue-700 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:shadow-lg transition duration-150 ease-in-out">Prev Page</button>
            </nuxt-link>
            <nuxt-link v-if="nextPage" :to="`/page/${pageNumber + 1}`">
              <button type="button" class="inline-block px-6 py-2.5 bg-blue-500 text-white font-medium text-xs leading-tight uppercase rounded shadow-md hover:bg-blue-700 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:shadow-lg transition duration-150 ease-in-out">Next Page</button>
            </nuxt-link>
          </section> -->
        </div>
      </div>
      <Footer/>
    </div>
  </div>
</template>

<script>
import Header from '../../components/Header.vue'
import Footer from '../../components/Footer.vue'
import ExpeditionPreview from '../../components/ExpeditionPreview.vue'

export default {
  data() {
    return {
      travels: {},
      title: "",
      nextPage: false,
      pageNumber: 0,
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
    ExpeditionPreview,
  },
  async asyncData({ params, $axios }) {
    const pageNumber = parseInt(params.number)
    const travels = await $axios.$get(process.env.baseAPI + '/travels?limit=10&offset=' + pageNumber)
    
    const nextPage = travels.length === 10
    return { nextPage, travels, pageNumber }
  },
  computed: {
    prevLink() {
      return this.pageNumber === 2 ? '/' : `/page/${this.pageNumber - 1}`
    }
  },

}
</script>