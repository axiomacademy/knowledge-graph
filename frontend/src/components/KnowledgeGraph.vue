<template>
  <div class="relative">
    <div class="flex flex-col items-center">
      <svg id="knowledge-graph" width="1000" height="300"></svg>
    </div>
    <button class="primary-color w-10 h-10 text-white rounded text-2xl shadow ripple hover:shadow-lg focus:outline-none absolute right-8 bottom-2">+</button>
  </div>
  <hr class="my-4 mx-8">
</template>

<script>
import * as d3 from 'd3'

export default {
  name: 'KnowledgeGraph',
  props: {
    concepts: Array
  },
  data() {
    return {
      test: [{
        uuid: "1",
        title: "Hello world",
        prerequisites: []
      },{
        uuid: "2",
        title: "Hello world 2",
        prerequisites: ["1"],
      },{
        uuid: "3",
        title: "Hello world 3",
        prerequisites: ["1"],
      },{
        uuid: "4",
        title: "Hello world 4",
        prerequisites: ["2", "3"]
      }]
    }
  },
  computed: {
    nodes() {
      return this.concepts.map((elem) => ({id: elem.uuid, title: elem.title}))
    },
    edges() {
      let edges = []
      for (const elem of this.concepts) {
        if(elem.prerequisites == null) {
          continue
        }

        let elemEdges = elem.prerequisites.map((prereq) => ({source: prereq, target: elem.uuid}))
        console.log(elemEdges)
        edges.push(...elemEdges)
      }
      return edges
    }
  },
  mounted() {
    console.log(this.concepts)
    console.log(this.nodes)
    console.log(this.edges)
    this.constructGraph()
  },
  methods: {
    constructGraph() {
      return new Promise((resolve,) => {
        let svg = d3.select('svg')
        let width = +svg.attr('width')
        let height = +svg.attr('height')

        let color = d3.scaleOrdinal() // D3 Version 4
          .domain(['Black', 'White'])
          .range(['#000000', '#FFFFFF'])
        
        svg.attr("align","center")

        svg.append('defs').append('marker')
          .attr('id', 'arrowhead')
          .attr('viewBox', '-0 -5 10 10')
          .attr('refX', 13)
          .attr('refY', 0)
          .attr('orient', 'auto')
          .attr('markerWidth', 5)
          .attr('markerHeight', 8)
          .attr('xoverflow', 'visible')
          .append('svg:path')
          .attr('d', 'M 0,-5 L 10 ,0 L 0,5')
          .attr('fill', 'Black')
          .style('stroke','none')

        let simulation = d3.forceSimulation()
          .force('link', d3.forceLink().id((d) => d.id ).distance(200))
          .force('charge', d3.forceManyBody().strength(-300))
          .force('center', d3.forceCenter(width / 2, height / 2))
          .force('collision', d3.forceCollide().radius(10))
        
        let link = svg.append('g')
          .attr('class', 'links')
          .selectAll('line')
          .data(this.edges)
          .enter().append('line')
          .attr('stroke', color('Black'))
          .attr('stroke-width', 2)
          .attr('marker-end','url(#arrowhead)')

        let node = svg.append('g')
          .attr('class', 'nodes')
          .selectAll('g')
          .data(this.nodes)
          .enter().append('g')

        // The actual nodes
        node.append('circle')
          .attr('r', 10)
          .attr('fill', '#7938D8')
          .call(d3.drag()
              .on('start', dragstarted)
              .on('drag', dragged)
             )
          .on('click', onclick.bind(this))

        // The labels
        node.append('text')
            .text(function(d) {
              return d.title
            })
            .attr('text-anchor', 'middle')
            .attr('transform', 'translate(0, -30)')
            .attr('fill', () => color('Black'))
            .attr('font-size', 15)
            .attr('font-family', 'Roboto')
            .attr('font-weight', 300)
            .attr('stroke', color('Black'))

        node.append('title')
            .text(function(d) { return d.title })

        simulation
            .nodes(this.nodes)
            .on('tick', ticked)

        simulation.force('link')
            .links(this.edges) 

        function ticked() {
          link
              .attr('x1', function(d) { return d.source.x })
              .attr('y1', function(d) { return d.source.y })
              .attr('x2', function(d) { return d.target.x })
              .attr('y2', function(d) { return d.target.y })

          node
              .attr('transform', function(d) {
                return 'translate(' + d.x + ',' + d.y + ')'
              })
        }

        function dragstarted(event, d) {
          if (!event.active) simulation.alphaTarget(0.3).restart()
          d.fx = d.x
          d.fy = d.y
        }

        function dragged(event, d) {
          d.fx = event.x
          d.fy = event.y
        }

        /*function dragended(event, d) {
          if (!event.active) simulation.alphaTarget(0)
          d.fx = null
          d.fy = null
        }*/

        function onclick(d) {
          this.$emit('conceptClicked', d.srcElement.__data__.id)
          resolve(d)
        }
      })
    }  
  }
}
</script>

<style scoped>

</style>
