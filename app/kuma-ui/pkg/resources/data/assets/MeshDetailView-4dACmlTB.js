import{d as D,l as M,o as p,c as N,f as o,t as a,p as e,e as t,P as x,_ as T,S as E,a as c,b as v,w as r,m as w,B as V,q as I}from"./index-CImj3nNu.js";import{_ as $}from"./ResourceCodeBlock.vue_vue_type_script_setup_true_lang-DbAPv0m9.js";import"./CodeBlock-CsNe8TyX.js";import"./toYaml-DB9FPXFY.js";const A={class:"date-status"},q=D({__name:"ResourceDateStatus",props:{creationTime:{},modificationTime:{}},setup(l){const{t:s}=M(),n=l;return(i,S)=>(p(),N("span",A,[o(a(e(s)("common.detail.created"))+": "+a(e(s)("common.formats.datetime",{value:Date.parse(n.creationTime)}))+" "+a()+" ",1),t(e(x)),o(" "+a(e(s)("common.detail.modified"))+":"+a(e(s)("common.formats.datetime",{value:Date.parse(n.modificationTime)})),1)]))}}),y=T(q,[["__scopeId","data-v-785cac69"]]),P={class:"stack"},j={class:"date-status-wrapper"},z=D({__name:"MeshDetailView",props:{mesh:{}},setup(l){const s=l,n=E();return(i,S)=>{const g=c("RouteTitle"),_=c("DataSource"),C=c("AppView"),k=c("RouteView");return p(),v(k,{name:"mesh-detail-view",params:{mesh:""}},{default:r(({route:d,t:R,uri:h})=>[t(g,{title:R("meshes.routes.overview.title"),render:!1},null,8,["title"]),o(),t(C,null,{default:r(()=>[w("div",P,[t(_,{src:h(e(V),"/mesh-insights/:name",{name:d.params.mesh})},{default:r(({data:m})=>[t(e(n),{mesh:s.mesh,"mesh-insight":m},null,8,["mesh","mesh-insight"])]),_:2},1032,["src"]),o(),t($,{resource:i.mesh.config},{default:r(({copy:m,copying:B})=>[B?(p(),v(_,{key:0,src:h(e(V),"/meshes/:name/as/kubernetes",{name:d.params.mesh},{cacheControl:"no-store"}),onChange:u=>{m(f=>f(u))},onError:u=>{m((f,b)=>b(u))}},null,8,["src","onChange","onError"])):I("",!0)]),_:2},1032,["resource"]),o(),w("div",j,[t(y,{"creation-time":i.mesh.creationTime,"modification-time":i.mesh.modificationTime},null,8,["creation-time","modification-time"])])])]),_:2},1024)]),_:1})}}}),K=T(z,[["__scopeId","data-v-3b0f5e65"]]);export{K as default};
