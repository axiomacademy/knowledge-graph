<template>
  <v-app>
    <v-main>
      <h1>Hello</h1>
      <KnowledgeGraph :width="1.0" :height="1.0" v-if="dataReady" :graph="this.g" v-on:concept-clicked="selectedConcept"/>
    </v-main>
  </v-app>
</template>

<script>
// @ is an alias to /src
import KnowledgeGraph from '@/components/KnowledgeGraph.vue'

import { getAllConcepts } from '../services/ConceptService'

import * as dagreD3 from 'dagre-d3'

export default {
  name: 'Home',
  data() {
    return {
      g: null,
      dataReady: false,
      selectedConcept: null,
    }
  },
  components: {
    KnowledgeGraph,
  },
  async mounted() {
    await this.getConcepts()
  },
  methods: {
    async getConcepts() {
      this.dataReady = false
      
      const res = await getAllConcepts()

      let graph = new dagreD3.graphlib.Graph()
      graph.setGraph({rankdir: "LR"})
      graph.setDefaultEdgeLabel(function() { return {} })

      if(res.concepts != null) {
        for(let i=0; i < res.concepts.length; i++) {
          graph.setNode(res.concepts[i].uuid, { label: res.concepts[i].title, shape: "circle", content: res.concepts[i].content }) 
        }
      }
      
      if(res.links != null) {
        for(let i=0; i < res.links.length; i++) {
          graph.setEdge(res.links[i].start_id, res.links[i].end_id) 
        }
      }

      this.g = graph
      this.dataReady = true
    },
    selectConcept(uuid) {
      const node = this.g.node(uuid)
      this.selectedConcept = {
        "uuid": uuid,
        "title": node.label,
        "content": node.content
      }
    },
  }
}
</script>
