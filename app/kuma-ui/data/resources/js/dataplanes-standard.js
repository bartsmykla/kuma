(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["dataplanes-standard"],{"4a34":function(t,e,a){"use strict";a.d(e,"a",(function(){return l}));a("5db7"),a("a630"),a("d81d"),a("73d9"),a("4fad"),a("d3b7"),a("ac1f"),a("6062"),a("3ca3"),a("1276"),a("ddb0");function n(t){if(Array.isArray(t))return t}a("a4d3"),a("e01a"),a("d28b");function r(t,e){if("undefined"!==typeof Symbol&&Symbol.iterator in Object(t)){var a=[],n=!0,r=!1,i=void 0;try{for(var s,o=t[Symbol.iterator]();!(n=(s=o.next()).done);n=!0)if(a.push(s.value),e&&a.length===e)break}catch(l){r=!0,i=l}finally{try{n||null==o["return"]||o["return"]()}finally{if(r)throw i}}return a}}var i=a("06c5");function s(){throw new TypeError("Invalid attempt to destructure non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")}function o(t,e){return n(t)||r(t,e)||Object(i["a"])(t,e)||s()}function l(t){var e=[],a=t.networking.inbound||null;a&&(e=a.flatMap((function(t){return Object.entries(t.tags)})).map((function(t){var e=o(t,2),a=e[0],n=e[1];return a+"="+n})));var n=t.networking.gateway||null;return n&&(e=Object.entries(n.tags).map((function(t){var e=o(t,2),a=e[0],n=e[1];return a+"="+n}))),e=Array.from(new Set(e)),e.map((function(t){return t.split("=")})).map((function(t){var e=o(t,2),a=e[0],n=e[1];return{label:a,value:n}}))}},"55ca":function(t,e,a){"use strict";var n=a("8af9"),r=a.n(n);r.a},"5db7":function(t,e,a){"use strict";var n=a("23e7"),r=a("a2bf"),i=a("7b0b"),s=a("50c4"),o=a("1c0b"),l=a("65f0");n({target:"Array",proto:!0},{flatMap:function(t){var e,a=i(this),n=s(a.length);return o(t),e=l(a,0),e.length=r(e,a,a,n,0,1,t,arguments.length>1?arguments[1]:void 0),e}})},6062:function(t,e,a){"use strict";var n=a("6d61"),r=a("6566");t.exports=n("Set",(function(t){return function(){return t(this,arguments.length?arguments[0]:void 0)}}),r)},6566:function(t,e,a){"use strict";var n=a("9bf2").f,r=a("7c73"),i=a("e2cc"),s=a("0366"),o=a("19aa"),l=a("2266"),u=a("7dd0"),c=a("2626"),p=a("83ab"),f=a("f183").fastKey,d=a("69f3"),v=d.set,h=d.getterFor;t.exports={getConstructor:function(t,e,a,u){var c=t((function(t,n){o(t,c,e),v(t,{type:e,index:r(null),first:void 0,last:void 0,size:0}),p||(t.size=0),void 0!=n&&l(n,t[u],t,a)})),d=h(e),y=function(t,e,a){var n,r,i=d(t),s=m(t,e);return s?s.value=a:(i.last=s={index:r=f(e,!0),key:e,value:a,previous:n=i.last,next:void 0,removed:!1},i.first||(i.first=s),n&&(n.next=s),p?i.size++:t.size++,"F"!==r&&(i.index[r]=s)),t},m=function(t,e){var a,n=d(t),r=f(e);if("F"!==r)return n.index[r];for(a=n.first;a;a=a.next)if(a.key==e)return a};return i(c.prototype,{clear:function(){var t=this,e=d(t),a=e.index,n=e.first;while(n)n.removed=!0,n.previous&&(n.previous=n.previous.next=void 0),delete a[n.index],n=n.next;e.first=e.last=void 0,p?e.size=0:t.size=0},delete:function(t){var e=this,a=d(e),n=m(e,t);if(n){var r=n.next,i=n.previous;delete a.index[n.index],n.removed=!0,i&&(i.next=r),r&&(r.previous=i),a.first==n&&(a.first=r),a.last==n&&(a.last=i),p?a.size--:e.size--}return!!n},forEach:function(t){var e,a=d(this),n=s(t,arguments.length>1?arguments[1]:void 0,3);while(e=e?e.next:a.first){n(e.value,e.key,this);while(e&&e.removed)e=e.previous}},has:function(t){return!!m(this,t)}}),i(c.prototype,a?{get:function(t){var e=m(this,t);return e&&e.value},set:function(t,e){return y(this,0===t?0:t,e)}}:{add:function(t){return y(this,t=0===t?0:t,t)}}),p&&n(c.prototype,"size",{get:function(){return d(this).size}}),c},setStrong:function(t,e,a){var n=e+" Iterator",r=h(e),i=h(n);u(t,e,(function(t,e){v(this,{type:n,target:t,state:r(t),kind:e,last:void 0})}),(function(){var t=i(this),e=t.kind,a=t.last;while(a&&a.removed)a=a.previous;return t.target&&(t.last=a=a?a.next:t.state.first)?"keys"==e?{value:a.key,done:!1}:"values"==e?{value:a.value,done:!1}:{value:[a.key,a.value],done:!1}:(t.target=void 0,{value:void 0,done:!0})}),a?"entries":"values",!a,!0),c(e)}}},6663:function(t,e,a){"use strict";var n=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"entity-url-control"},[t.shouldDisplay?a("KClipboardProvider",{scopedSlots:t._u([{key:"default",fn:function(e){var n=e.copyToClipboard;return[a("KPop",{attrs:{placement:"bottom"}},[a("KButton",{attrs:{appearance:"secondary",size:"small"},on:{click:function(){n(t.url)}}},[a("KIcon",{attrs:{slot:"icon",icon:"externalLink"},slot:"icon"}),t._v(" "+t._s(t.copyButtonText)+" ")],1),a("div",{attrs:{slot:"content"},slot:"content"},[a("p",[t._v(t._s(t.confirmationText))])])],1)]}}],null,!1,1603401634)}):t._e()],1)},r=[],i={name:"EntityURLControl",props:{url:{type:String,required:!0},copyButtonText:{type:String,default:"Copy URL"},confirmationText:{type:String,default:"URL copied to clipboard!"}},computed:{shouldDisplay:function(){var t=this.$route.params.mesh||null;return!(!t||"all"===t)}}},s=i,o=a("2877"),l=Object(o["a"])(s,n,r,!1,null,null,null);e["a"]=l.exports},"6d61":function(t,e,a){"use strict";var n=a("23e7"),r=a("da84"),i=a("94ca"),s=a("6eeb"),o=a("f183"),l=a("2266"),u=a("19aa"),c=a("861d"),p=a("d039"),f=a("1c7e"),d=a("d44e"),v=a("7156");t.exports=function(t,e,a){var h=-1!==t.indexOf("Map"),y=-1!==t.indexOf("Weak"),m=h?"set":"add",b=r[t],g=b&&b.prototype,w=b,x={},E=function(t){var e=g[t];s(g,t,"add"==t?function(t){return e.call(this,0===t?0:t),this}:"delete"==t?function(t){return!(y&&!c(t))&&e.call(this,0===t?0:t)}:"get"==t?function(t){return y&&!c(t)?void 0:e.call(this,0===t?0:t)}:"has"==t?function(t){return!(y&&!c(t))&&e.call(this,0===t?0:t)}:function(t,a){return e.call(this,0===t?0:t,a),this})};if(i(t,"function"!=typeof b||!(y||g.forEach&&!p((function(){(new b).entries().next()})))))w=a.getConstructor(e,t,h,m),o.REQUIRED=!0;else if(i(t,!0)){var D=new w,k=D[m](y?{}:-0,1)!=D,O=p((function(){D.has(1)})),T=f((function(t){new b(t)})),_=!y&&p((function(){var t=new b,e=5;while(e--)t[m](e,e);return!t.has(-0)}));T||(w=e((function(e,a){u(e,w,t);var n=v(new b,e,w);return void 0!=a&&l(a,n[m],n,h),n})),w.prototype=g,g.constructor=w),(O||_)&&(E("delete"),E("has"),h&&E("get")),(_||k)&&E(m),y&&g.clear&&delete g.clear}return x[t]=w,n({global:!0,forced:w!=b},x),d(w,t),y||a.setStrong(w,t,h),w}},"720e":function(t,e,a){"use strict";a.r(e);var n=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"standard-dataplanes"},[a("FrameSkeleton",[a("DataOverview",{attrs:{"page-size":t.pageSize,"has-error":t.hasError,"is-loading":t.isLoading,"empty-state":t.empty_state,"display-data-table":!0,"table-data":t.tableData,"table-data-is-empty":t.tableDataIsEmpty,"table-data-function-text":"View","table-data-row":"name"},on:{tableAction:t.tableAction,reloadData:t.loadData}},[a("template",{slot:"additionalControls"},[a("KButton",{staticClass:"add-dp-button",attrs:{appearance:"primary",size:"small",to:t.dataplaneWizardRoute}},[a("span",{staticClass:"custom-control-icon"},[t._v(" + ")]),t._v(" Create data plane proxy ")]),this.$route.query.ns?a("KButton",{staticClass:"back-button",attrs:{appearance:"primary",size:"small",to:{name:"dataplanes"}}},[a("span",{staticClass:"custom-control-icon"},[t._v(" ← ")]),t._v(" View All ")]):t._e()],1),a("template",{slot:"pagination"},[a("Pagination",{attrs:{"has-previous":t.previous.length>0,"has-next":t.hasNext},on:{next:t.goToNextPage,previous:t.goToPreviousPage}})],1)],2),!1===t.isEmpty?a("Tabs",{attrs:{"has-error":t.hasError,"is-loading":t.isLoading,tabs:t.tabs,"initial-tab-override":"overview"}},[a("template",{slot:"tabHeader"},[a("div",[a("h3",[t._v(t._s(t.tabGroupTitle))])]),a("div",[a("EntityURLControl",{attrs:{url:t.shareUrl}})],1)]),a("template",{slot:"overview"},[a("LabelList",{attrs:{"has-error":t.entityHasError,"is-loading":t.entityIsLoading,"is-empty":t.entityIsEmpty}},[a("div",[a("ul",t._l(t.entity.basicData,(function(e,n){return a("li",{key:n},[a("h4",[t._v(t._s(n))]),a("p",[t._v(" "+t._s(e)+" ")])])})),0)]),a("div",[a("h4",[t._v("Tags")]),a("p",t._l(t.entity.tags,(function(e,n){return a("span",{key:n,staticClass:"tag-cols"},[a("span",[t._v(" "+t._s(e.label)+": ")]),a("span",[t._v(" "+t._s(e.value)+" ")])])})),0)])])],1),a("template",{slot:"mtls"},[a("LabelList",{attrs:{"has-error":t.entityHasError,"is-loading":t.entityIsLoading,"is-empty":t.entityIsEmpty}},[t.entity.mtls?a("ul",t._l(t.entity.mtls,(function(e,n){return a("li",{key:n},[a("h4",[t._v(t._s(e.label))]),a("p",[t._v(" "+t._s(e.value)+" ")])])})),0):a("KAlert",{attrs:{appearance:"danger"}},[a("template",{slot:"alertMessage"},[t._v(" This data plane proxy does not yet have mTLS configured — "),a("a",{staticClass:"external-link",attrs:{href:"https://kuma.io/docs/"+t.version+"/documentation/security/#certificates",target:"_blank"}},[t._v(" Learn About Certificates in Kuma ")])])],2)],1)],1),a("template",{slot:"yaml"},[a("YamlView",{attrs:{title:t.entityOverviewTitle,"has-error":t.entityHasError,"is-loading":t.entityIsLoading,"is-empty":t.entityIsEmpty,content:t.rawEntity}})],1)],2):t._e()],1)],1)},r=[],i=(a("99af"),a("4160"),a("13d5"),a("b0c0"),a("d3b7"),a("159b"),a("96cf"),a("1da1")),s=a("5530"),o=a("2f62"),l=a("d7c2"),u=a("6663"),c=a("8218"),p=a("1d10"),f=a("1799"),d=a("2778"),v=a("251b"),h=a("ff9d"),y=a("0ada"),m=a("4a34"),b={name:"StandardDataplanes",metaInfo:{title:"Standard Data plane proxies"},components:{EntityURLControl:u["a"],FrameSkeleton:p["a"],Pagination:f["a"],DataOverview:d["a"],Tabs:v["a"],YamlView:h["a"],LabelList:y["a"]},mixins:[c["a"]],data:function(){return{isLoading:!0,isEmpty:!1,hasError:!1,entityIsLoading:!0,entityIsEmpty:!1,entityHasError:!1,tableDataIsEmpty:!1,empty_state:{title:"No Data",message:"There are no Standard data plane proxies present."},tableData:{headers:[{key:"actions",hideLabel:!0},{label:"Status",key:"status"},{label:"Name",key:"name"},{label:"Mesh",key:"mesh"},{label:"Type",key:"type"},{label:"Tags",key:"tags"},{label:"Last Connected",key:"lastConnected"},{label:"Last Updated",key:"lastUpdated"},{label:"Total Updates",key:"totalUpdates"},{label:"Kuma DP version",key:"dpVersion"},{label:"Envoy version",key:"envoyVersion"}],data:[]},tabs:[{hash:"#overview",title:"Overview"},{hash:"#mtls",title:"Certificate Insights"},{hash:"#yaml",title:"YAML"}],entity:[],rawEntity:null,firstEntity:null,pageSize:this.$pageSize,pageOffset:null,next:null,hasNext:!1,previous:[],tabGroupTitle:null,entityNamespace:null,entityOverviewTitle:null,showmTLSTab:!1}},computed:Object(s["a"])(Object(s["a"])({},Object(o["b"])({environment:"getEnvironment",queryNamespace:"getItemQueryNamespace"})),{},{dataplaneWizardRoute:function(){return"universal"===this.environment?{name:"universal-dataplane"}:{name:"kubernetes-dataplane"}},version:function(){var t=this.$store.getters.getVersion;return null!==t?t:"latest"},shareUrl:function(){var t=this,e="".concat(window.location.origin,"/#"),a=this.entity,n=function(){return a.basicData?t.$route.query.ns?t.$route.fullPath:"".concat(e).concat(t.$route.fullPath,"?ns=").concat(a.basicData.name):null};return n()}}),watch:{$route:function(t,e){this.loadData()}},beforeMount:function(){this.loadData()},methods:{init:function(){this.loadData()},goToPreviousPage:function(){this.pageOffset=this.previous.pop(),this.next=null,this.loadData()},goToNextPage:function(){this.previous.push(this.pageOffset),this.pageOffset=this.next,this.next=null,this.loadData()},tableAction:function(t){var e=t;this.$store.dispatch("updateSelectedTab",this.tabs[0].hash),this.$store.dispatch("updateSelectedTableRow",e.name),this.getEntity(e)},loadData:function(){var t=this;this.isLoading=!0;var e=this.$route.params.mesh||null,a=this.$route.query.ns||null,n={size:this.pageSize,offset:this.pageOffset,gateway:!1,ingress:!1},r=function(){return"all"===e?t.$api.getAllDataplaneOverviews(n):a&&a.length&&"all"!==e?t.$api.getDataplaneOverviewFromMesh(e,a):t.$api.getAllDataplaneOverviewsFromMesh(e,n)},i=function(e,a,n){t.$api.getDataplaneOverviewFromMesh(e,a).then((function(e){var a,r,i="n/a",s=[],o=[],u=[],c="Offline",p="",f="",d=[],v=[],h=(e.dataplane.networking.inbound,"Standard");if(s=Object(m["a"])(e.dataplane),e.dataplaneInsight.subscriptions&&e.dataplaneInsight.subscriptions.length){e.dataplaneInsight.subscriptions.forEach((function(t){var e=t.status.total.responsesSent||0,a=t.status.total.responsesRejected||0,n=t.connectTime||i,r=t.status.lastUpdateTime||i,s=t.disconnectTime||null;o.push(parseInt(e)),u.push(parseInt(a)),d.push(n),v.push(r),c=n&&n.length&&!s?"Online":"Offline",t.version&&t.version.kumaDp&&(p=t.version.kumaDp.version,f=t.version.envoy.version)})),o=o.reduce((function(t,e){return t+e})),u=u.reduce((function(t,e){return t+e}));var y=d.reduce((function(t,e){return t&&e?t.MeasureDate>e.MeasureDate?t:e:null})),b=v.reduce((function(t,e){return t&&e?t.MeasureDate>e.MeasureDate?t:e:null})),g=new Date(y),w=new Date(b);a=y&&!isNaN(g)?Object(l["d"])(g):"never",r=b&&!isNaN(w)?Object(l["d"])(w):"never"}else a="never",r="never",o=0,u=0,p="-",f="-";return n.push({name:e.name,mesh:e.mesh,tags:s,status:c,lastConnected:a,lastUpdated:r,totalUpdates:o,totalRejectedUpdates:u,dpVersion:p,envoyVersion:f,type:h}),t.sortEntities(n),n})).catch((function(t){console.error(t)}))},s=function(){return r().then((function(n){var r=function(){var e=n;return"total"in e?0!==e.total&&e.items&&e.items.length>0?t.sortEntities(e.items):null:e};if(r()){n.next?(t.next=Object(l["b"])(n.next),t.hasNext=!0):t.hasNext=!1;var s=[],o=a?r():r()[0];t.firstEntity=o.name,t.getEntity(o),t.$store.dispatch("updateSelectedTableRow",t.firstEntity),a&&a.length&&e&&e.length?i(e,a,s):r().forEach((function(t){i(t.mesh,t.name,s)})),t.tableData.data=s,t.tableDataIsEmpty=!1,t.isEmpty=!1}else t.tableData.data=[],t.tableDataIsEmpty=!0,t.isEmpty=!0,t.getEntity(null)})).catch((function(e){t.hasError=!0,t.isEmpty=!0,console.error(e)})).finally((function(){setTimeout((function(){t.isLoading=!1}),"500")}))};s()},getEntity:function(t){var e=this;this.entityIsLoading=!0,this.entityIsEmpty=!1;var a=this.$route.params.mesh;if(t&&null!==t){var n="all"===a?t.mesh:a;return this.$api.getDataplaneFromMesh(n,t.name).then((function(a){if(a){var r=["type","name","mesh"],o=function(){var a=Object(i["a"])(regeneratorRuntime.mark((function a(){var r,i,s,o,u,c;return regeneratorRuntime.wrap((function(a){while(1)switch(a.prev=a.next){case 0:return r=null,a.prev=1,a.next=4,e.$api.getDataplaneOverviewFromMesh(n,t.name);case 4:i=a.sent,i.dataplaneInsight.mTLS&&(s=i.dataplaneInsight.mTLS,o=new Date(s.certificateExpirationTime),u=new Date(o.getTime()+6e4*o.getTimezoneOffset()),c="\n                      ".concat(u.toLocaleDateString("en-US")," ").concat(u.getHours(),":").concat(u.getMinutes(),":").concat(u.getSeconds(),"\n                    "),r={certificateExpirationTime:{label:"Expiration Time",value:c},lastCertificateRegeneration:{label:"Last Generated",value:Object(l["d"])(s.lastCertificateRegeneration)},certificateRegenerations:{label:"Regenerations",value:s.certificateRegenerations}}),a.next=11;break;case 8:a.prev=8,a.t0=a["catch"](1),console.error(a.t0);case 11:return a.abrupt("return",r);case 12:case"end":return a.stop()}}),a,null,[[1,8]])})));return function(){return a.apply(this,arguments)}}(),u=function(){var t=Object(i["a"])(regeneratorRuntime.mark((function t(){return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return t.t0=Object(s["a"])({},Object(l["c"])(a,r)),t.t1=Object(m["a"])(a),t.next=4,o();case 4:return t.t2=t.sent,t.abrupt("return",{basicData:t.t0,tags:t.t1,mtls:t.t2});case 6:case"end":return t.stop()}}),t)})));return function(){return t.apply(this,arguments)}}();u().then((function(t){e.entity=t,e.entityNamespace=t.basicData.name,e.tabGroupTitle="Mesh: ".concat(t.basicData.name),e.entityOverviewTitle="Entity Overview for ".concat(t.basicData.name)})),e.rawEntity=Object(l["h"])(a)}else e.entity=null,e.entityIsEmpty=!0})).catch((function(t){e.entityHasError=!0,console.error(t)})).finally((function(){setTimeout((function(){e.entityIsLoading=!1}),"500")}))}setTimeout((function(){e.entityIsEmpty=!0,e.entityIsLoading=!1}),"500")}}},g=b,w=(a("55ca"),a("2877")),x=Object(w["a"])(g,n,r,!1,null,"5ca9061d",null);e["default"]=x.exports},"73d9":function(t,e,a){var n=a("44d2");n("flatMap")},"8af9":function(t,e,a){},a2bf:function(t,e,a){"use strict";var n=a("e8b5"),r=a("50c4"),i=a("0366"),s=function(t,e,a,o,l,u,c,p){var f,d=l,v=0,h=!!c&&i(c,p,3);while(v<o){if(v in a){if(f=h?h(a[v],v,e):a[v],u>0&&n(f))d=s(t,e,f,r(f.length),d,u-1)-1;else{if(d>=9007199254740991)throw TypeError("Exceed the acceptable array length");t[d]=f}d++}v++}return d};t.exports=s},bb2f:function(t,e,a){var n=a("d039");t.exports=!n((function(){return Object.isExtensible(Object.preventExtensions({}))}))},f183:function(t,e,a){var n=a("d012"),r=a("861d"),i=a("5135"),s=a("9bf2").f,o=a("90e3"),l=a("bb2f"),u=o("meta"),c=0,p=Object.isExtensible||function(){return!0},f=function(t){s(t,u,{value:{objectID:"O"+ ++c,weakData:{}}})},d=function(t,e){if(!r(t))return"symbol"==typeof t?t:("string"==typeof t?"S":"P")+t;if(!i(t,u)){if(!p(t))return"F";if(!e)return"E";f(t)}return t[u].objectID},v=function(t,e){if(!i(t,u)){if(!p(t))return!0;if(!e)return!1;f(t)}return t[u].weakData},h=function(t){return l&&y.REQUIRED&&p(t)&&!i(t,u)&&f(t),t},y=t.exports={REQUIRED:!1,fastKey:d,getWeakData:v,onFreeze:h};n[u]=!0}}]);