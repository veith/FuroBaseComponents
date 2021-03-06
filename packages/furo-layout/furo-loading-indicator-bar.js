import {LitElement, html, css} from 'lit-element';
import {FBP} from "@furo/fbp";
import {Theme} from "@furo/framework/theme"

/**
 * `loading-indicator`
 * Describe your element
 * todo: make it like https://material.io/design/components/progress-indicators.html#
 *
 * @summary shortdescription
 * @customElement
 * @demo demo/loading-indicator.html
 * @appliesMixin FBP
 */
class FuroLoadingIndicatorBar extends FBP(LitElement) {

  constructor() {
    super();
    this.setAttribute("hidden", "")
  }


  /**
   *
   * @private
   * @return {CSSResult}
   */
  static get styles() {
    // language=CSS
    return Theme.getThemeForComponent(this.name) || css`
        :host {
            display: block;
            width: 100%;
            height: 4px;
            pointer-events: none;
            opacity: 0.8;
            background-color: var(--secondary-color);
            animation: colorchange 5s infinite;
            @apply --loading-indicator-mixin;
        }

        :host([hidden]) {
            display: none;
        }

        @keyframes colorchange {
            0% {
                background: red;
            }
            25% {
                background: yellow;
            }
            50% {
                background: blue;
            }
            75% {
                background: green;
            }
            100% {
                background: red;
            }
        }
    `
  }


  start() {
    this.removeAttribute("hidden");
  }

  stop() {
    this.setAttribute("hidden", "");
  }

  /**
   * @private
   * @returns {TemplateResult}
   */
  render() {
    // language=HTML
    return html`
      <slot></slot>
    `;
  }
}

window.customElements.define('furo-loading-indicator-bar', FuroLoadingIndicatorBar);
