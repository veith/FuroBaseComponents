import {LitElement, html, css} from 'lit-element';
import {Theme} from "@furo/framework/theme"
import "@furo/layout/furo-horizontal-flex"

/**
 * `furo-button-bar`
 *
 *
 * @summary
 * @customElement
 * @demo demo/furo-button-bar.html
 */
class FuroButtonBar extends (LitElement) {

    bindEntity(entity){
        // todo bind entity to do some state management
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
        }

        ::slotted(*) {
            margin: var(--spacing-xs, 8px);
        }

        ::slotted(*:first-child) {
            margin-left: 0;
        }

        ::slotted(*:last-child) {
            margin-right: 0;
        }

        furo-horizontal-flex{
          flex-wrap: wrap;
        }

    `
  }

  render() {
    // language=HTML
    return html`        
        <furo-horizontal-flex>
          <slot></slot>
        </furo-horizontal-flex>
    `;
  }

}

window.customElements.define('furo-button-bar', FuroButtonBar);
