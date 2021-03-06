import {LitElement, html, css} from 'lit-element';
import {Theme} from "@furo/framework/theme"
import {FBP} from "@furo/fbp";
import "@furo/util"
import "@furo/layout"
import "@furo/doc-helper"

/**
 * `panel-doc`
 *  i didnt use furo-doc-panel because the paths are completly different in the page i.e /api/PACKAGENAME/doc
 *
 * @summary todo shortdescription
 * @customElement
 * @demo demo/panel-doc.html
 * @appliesMixin FBP
 */
class PanelDoc extends FBP(LitElement) {

  /**
   * Themable Styles
   * @private
   * @return {CSSResult}
   */
  static get styles() {
    // language=CSS
    return Theme.getThemeForComponent(this.name) || css`
        :host {
            display: block;
            height: 100%;
            padding-top: var(--spacing);
            padding-left: var(--spacing-s);
            overflow: hidden;
            box-sizing: border-box;
            background-color: var(--surface);
            --split-master-width: 212px;
        }

        :host([hidden]) {
            display: none;
        }

        furo-doc-element, furo-doc-class {
            padding-left: var(--spacing);
        }

        /** the background of the bar itself. **/
        ::-webkit-scrollbar {
            width: 6px;
            background-color: var(--surface, white);
        }

        /** the directional buttons on the scrollbar. **/
        ::-webkit-scrollbar-button {
            background-color: var(--on-surface, black);
        }

        /** the empty space “below” the progress bar. **/
        ::-webkit-scrollbar-track {
        }

        /** the top-most layer of the the progress bar not covered by the thumb. **/
        ::-webkit-scrollbar-track-piece {
        }

        /** the draggable scrolling element resizes depending on the size of the scrollable element. **/
        ::-webkit-scrollbar-thumb {
            background-color: var(--on-surface, black);
            border-radius: 3px;
        }

        /** the bottom corner of the scrollable element, where two scrollbar meet. **/
        ::-webkit-scrollbar-corner {
        }

        /** the draggable resizing handle that appears above the scrollbar-corner at the bottom corner of some elements. **/
        ::-webkit-resizer {
        }
    `
  }

  /**
   * flow is ready lifecycle method
   */
  _FBPReady() {
    super._FBPReady();
    //this._FBPTraceWires()
    this._FBPAddWireHook("--packageChanged", e => {
      this._FBPTriggerWire("--src", "../../../node_modules/@furo/" + e.pathSegments[0] + "/analysis.json")
    })

  }


  /**
   * @private
   * @returns {TemplateResult}
   */
  render() {
    // language=HTML
    return html`

      <!-- find the package -->
      <furo-location url-space-regex="^/api" @-location-changed="--packageChanged"></furo-location>
      <furo-location url-space-regex="^/api/[^/]*/doc" @-location-changed="--pathChanged"></furo-location>
      <furo-doc-fetch-analysis ƒ-fetch-src="--src" ƒ-check-subroute="--pathChanged"
                               @-data="--analysis"></furo-doc-fetch-analysis>

      <furo-split-view>

        <!-- the doc menu -->
        <furo-doc-menu slot="master" scroll ƒ-analysis="--analysis" @-element="--element"
                       @-class="--class" @-mixin="--class"></furo-doc-menu>

        
        <furo-doc-element scroll ƒ-print="--element" ƒ-hide="--class"></furo-doc-element>


        <furo-doc-class scroll ƒ-print="--class" ƒ-hide="--element"></furo-doc-class>

      </furo-split-view>
    `;
  }
}

window.customElements.define('panel-doc', PanelDoc);
