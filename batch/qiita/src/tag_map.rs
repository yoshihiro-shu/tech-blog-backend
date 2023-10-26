use std::collections::HashMap;

pub fn create_map() -> HashMap<&'static str, &'static str> {
    let mut tag_map: HashMap<&str, &str> = HashMap::new();

    tag_map.insert("Docker", "docker");
    tag_map.insert("Kubernetes", "kubernetes");
    tag_map.insert("Golang", "golang");
    tag_map.insert("Agile", "agile");
    tag_map.insert("Requirement Definition", "requirement-definition");
    tag_map.insert("Nuxt", "nuxt");
    tag_map.insert("Network", "network");
    tag_map.insert("dns", "dns");
    // 日本語の部分を英語に変換
    tag_map.insert("インフラ", "infrastructure");
    tag_map.insert("アジャイル", "agile-methodology");
    tag_map.insert("プロジェクト管理", "project-management");
    tag_map.insert("チームビルディング", "team-building");
    tag_map.insert("ふりかえり", "reflection");
    tag_map.insert("プロジェクトマネジメント", "project-management");
    tag_map.insert("AI", "ai");
    tag_map.insert("ビジネス", "business");
    tag_map.insert("生産性向上", "productivity-improvement");
    tag_map.insert("Google", "google");
    tag_map.insert("マーケティング", "marketing");
    tag_map.insert("SEO対策", "seo-strategies");
    tag_map.insert("解決", "problem-solving");
    tag_map.insert("論理的思考", "logical-thinking");
    tag_map.insert("リーダー", "leader");
    tag_map.insert("kubernetes", "kubernetes");
    tag_map.insert("kubectl", "kubectl");
    tag_map.insert("ckad", "ckad");
    tag_map.insert("CKA", "cka");
    tag_map.insert("プレゼンテーション", "presentation");
    tag_map.insert("ロジカルシンキング", "logical-thinking");
    tag_map.insert("Go", "go");
    tag_map.insert("dockerfile", "dockerfile");
    tag_map.insert("DockerHub", "docker-hub");
    tag_map.insert("沼", "quagmire");
    tag_map.insert("個人開発", "personal-development");
    tag_map.insert("GitHubActions", "github-actions");
    tag_map.insert("要件定義", "requirement-definition");
    tag_map.insert("ユースケース", "use-case");
    tag_map.insert("ワイヤーフレーム", "wireframe");
    tag_map.insert("デザイン設計", "design-planning");
    tag_map.insert("Cloud", "cloud");
    tag_map.insert("docker-compose", "docker-compose");
    tag_map.insert("googlecloud", "googlecloud");
    tag_map.insert("cookie", "cookie");
    tag_map.insert("Vue.js", "vue-js");
    tag_map.insert("Vuex", "vuex");
    tag_map.insert("ssr", "ssr");
    tag_map.insert("開発環境", "development-environment");
    tag_map.insert("TypeScript", "typescript");
    tag_map.insert("Nuxt3", "nuxt3");

    tag_map
}
