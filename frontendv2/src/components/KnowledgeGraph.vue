<template>
  <div class="knowledge-graph">
    <div class="relative">
      <div class="flex flex-col items-center">
        <svg id="knowledge-graph" width="100%"><g></g></svg>
      </div>
      <button class="primary-color w-10 h-10 mb-2 text-white rounded text-2xl shadow ripple hover:shadow-lg focus:outline-none absolute right-8 bottom-2">+</button>
    </div>
    <hr class="mb-4 mx-8">
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
      svg.call(zoom.transform, d3.zoomIdentity.translate((svg.attr("width") - this.graph.graph().width * initialScale) / 2, 20).scale(initialScale))
    }
  }
}
</script>

<style>
.node rect,
.node circle,
.node ellipse {
  stroke: none;
  fill: var(--primary-color);
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
