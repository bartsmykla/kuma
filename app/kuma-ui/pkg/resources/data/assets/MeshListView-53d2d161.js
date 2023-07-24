import{q as u,J as h,T as d,l as b,a as f}from"./kongponents.es-21ce59a5.js";import{A as g,_ as w}from"./DataSource.vue_vue_type_script_setup_true_lang-c75e9562.js";import{i as y,A as v,_ as k}from"./RouteView.vue_vue_type_script_setup_true_lang-9e62d24f.js";import{_ as z}from"./RouteTitle.vue_vue_type_script_setup_true_lang-dcec85af.js";import{d as C,r as V,o as x,a as A,w as e,h as s,b as t,k as n,g as r,t as L}from"./index-bdbf5b57.js";const $={class:"stack"},R=C({__name:"MeshListView",props:{page:{},size:{}},setup(l){const o=l,{t:m}=y();return(N,B)=>{const p=V("RouterLink");return x(),A(k,{name:"mesh-list-view"},{default:e(({route:c})=>[s(w,{src:`/meshes?page=${o.page}&size=${o.size}`},{default:e(({data:a,error:_})=>[s(v,{breadcrumbs:[{to:{name:"mesh-list-view"},text:t(m)("meshes.routes.items.breadcrumbs")}]},{title:e(()=>[n("h1",null,[s(z,{title:t(m)("meshes.routes.items.title"),render:!0},null,8,["title"])])]),default:e(()=>[r(),n("div",$,[s(t(u),null,{body:e(()=>[s(g,{"data-testid":"mesh-collection","empty-state-title":t(m)("common.emptyState.title"),"empty-state-message":t(m)("common.emptyState.message",{type:"Meshes"}),headers:[{label:"Name",key:"name"},{label:"Actions",key:"actions",hideLabel:!0}],"page-number":o.page,"page-size":o.size,total:a==null?void 0:a.total,items:a==null?void 0:a.items,error:_,onChange:c.update},{name:e(({row:i})=>[s(p,{to:{name:"mesh-detail-view",params:{mesh:i.name}}},{default:e(()=>[r(L(i.name),1)]),_:2},1032,["to"])]),actions:e(({row:i})=>[s(t(h),{class:"actions-dropdown","kpop-attributes":{placement:"bottomEnd",popoverClasses:"mt-5 more-actions-popover"},width:"150"},{default:e(()=>[s(t(d),{class:"non-visual-button",appearance:"secondary",size:"small"},{icon:e(()=>[s(t(b),{color:"var(--black-400)",icon:"more",size:"16"})]),_:1})]),items:e(()=>[s(t(f),{item:{to:{name:"mesh-detail-view",params:{mesh:i.name}},label:"View"}},null,8,["item"])]),_:2},1024)]),_:2},1032,["empty-state-title","empty-state-message","page-number","page-size","total","items","error","onChange"])]),_:2},1024)])]),_:2},1032,["breadcrumbs"])]),_:2},1032,["src"])]),_:1})}}});export{R as default};