<template>
	<article>
		<div class="examples__block__map">
			<svg-map
				:map="World"
				location-role="link"
				@mouseover="pointLocation"
				@mouseout="unpointLocation"
				@focus="focusLocation"
        		@mousemove="moveOnLocation"
				@blur="blurLocation"
				@click="clickLocation"
			/>
			<div
				class="examples__block__map__tooltip"
				:style="tooltipStyle"
				>
				{{ pointedLocation }}
			</div>
		</div>
	</article>
</template>

<script>
import World from '@svg-maps/world'
import { getLocationName } from '../utilities'
import { SvgMap } from "vue-svg-map";
const URL = {
	ir: 'https://en.wikipedia.org/wiki/Auvergne-Rh%C3%B4ne-Alpes',
}
export default {
	name: 'Map',
	components: {
		SvgMap,
	},
	data() {
		return {
			World,
			pointedLocation: null,
			focusedLocation: null,
			clickedLocation: null,
      tooltipStyle: null,
		}
	},
	methods: {
		pointLocation(event) {
			this.pointedLocation = getLocationName(event.target)
		},
		unpointLocation(event) {
			this.pointedLocation = null
      this.tooltipStyle = { display: 'none' }
		},
		focusLocation(event) {
			this.focusedLocation = getLocationName(event.target)
		},
		blurLocation(event) {
			this.focusedLocation = null
		},
		clickLocation(event) {
			this.clickedLocation = getLocationName(event.target)
			window.open('/countries/' + event.target.id, '_blank')
		},
    moveOnLocation(event) {
			this.tooltipStyle = {
				display: 'block',
				top: `${event.clientY + 15}px`,
				left: `${event.clientX - 100}px`,
			}
		},
	},
}
</script>

<style lang="scss">
/* Using box model */
* {
  box-sizing: border-box;
}

/* Maps examples */
.examples {
	width: 1000px;
	margin: 0 auto;
	font-family: Arial, sans-serif;
	text-align: center;

	&__title {
		margin-bottom: 80px;
		padding-bottom: 40px;
		border-bottom: 1px dotted black;
		font-size: 24px;
	}

	&__block {
		margin-bottom: 80px;
		padding-bottom: 80px;
		border-bottom: 1px dotted black;


		&__title {
			margin-bottom: 40px;
			font-size: 20px;
		}

		&__info {
			display: inline-block;
			vertical-align: top;
			width: 450px;
			text-align: left;

			&__item {
				margin-bottom: 20px;
			}
		}

		&__map {
			display: inline-block;
			width: 100%;

			&--australia {
				.svg-map {
					stroke-width: 0.5; // Australia stroke is larger than other maps
				}
			}

			&--usa {
				width: 800px; // USA needs more space	for tooltip

				.svg-map__location {
					&--heat0 {
						fill: blue;
					}

					&--heat1 {
						fill: lightblue;
					}

					&--heat2 {
						fill: orange;
					}

					&--heat3 {
						fill: red;
					}

					&:focus,
					&:hover {
						opacity: 0.75;
					}
				}
			}

			&__tooltip {
				position: fixed;
				width: 200px;
				padding: 10px;
				border: 1px solid darkgray;
				background-color: white;
			}
		}
	}
}
</style>
