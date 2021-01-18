<template>
  <div class="knowledge-graph" style="position: relative;">
    <svg width="100%"><g></g></svg>
    <v-btn
      fab
      dark
      absolute
      color="primary"
      style="bottom: 32px; right: 32px;"
      @click="createNew"
      >
      <v-icon dark>
        mdi-plus
      </v-icon>
    </v-btn>
    <v-divider class="mx-8"></v-divider>
  </div>
</template>

<script>
import * as dagreD3 from 'dagre-d3'
import * as d3 from 'd3'

export default {
  name: 'KnowledgeGraph',
  props: {
    graph: Object
  },
  data() {
    return {
      w: window.innerWidth,
      h: window.innerHeight,
    }
  },
  mounted() {
    this.constructGraph()
  },
  methods: {
    createNew() {
      this.$emit("create-new")
    },
    constructGraph() {
      let svg = d3.select("svg"), inner = svg.select("g");
      
      // Set up zoom support
      let zoom = d3.zoom().on("zoom", function(event) {
        inner.attr("transform", event.transform)
      })
      svg.call(zoom)
      
      // Create the renderer
      let render = new dagreD3.render()
      render(inner, this.graph)
      
      // Set up click event
      inner.selectAll("g.node").on("click", function(event, d) {
        console.log("Hello")
        this.$emit('concept-clicked', d)
      }.bind(this))

      svg.attr("width", this.w)
      svg.attr("height", this.h * 0.8)  

      var initialScale = 1.0
      svg.call(zoom.transform, d3.zoomIdentity.translate((svg.attr("width") - this.graph.graph().width * initialScale) / 2,
        (svg.attr("height") - this.graph.graph().height * initialScale) / 2).scale(initialScale))
    }
  }
}
</script>

<style>
.node rect,
.node circle,
.node ellipse {
  stroke: none;
  fill: #7938D8;
  stroke-width: 1px;
  height: 50;
  width: 50;
}

.label text {
  font-family: 'IBM Plex Sans', sans-serif;
  font-size: 0.5rem;
  fill: white;
}

.edgePath path {
  stroke: #333;
  fill: #333;
  stroke-width: 1.5px;
}

</style>
