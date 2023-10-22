use std::collections::HashMap;

pub fn create_map() -> HashMap<&'static str, &'static str> {
    let mut name_to_slug: HashMap<&str, &str> = HashMap::new();

    name_to_slug.insert("Docker", "docker");
    name_to_slug.insert("Kubernetes", "kubernetes");
    name_to_slug.insert("Golang", "golang");
    name_to_slug.insert("Agile", "agile");
    name_to_slug.insert("Requirement definition", "requirement-definition");
    name_to_slug.insert("Nuxt", "nuxt");
    name_to_slug.insert("Network", "network");
    name_to_slug.insert("dns", "dns");
    // 日本語の部分を英語に変換
    name_to_slug.insert("インフラ", "infrastructure");
    name_to_slug.insert("アジャイル", "agile-methodology");
    name_to_slug.insert("プロジェクト管理", "project-management");
    name_to_slug.insert("チームビルディング", "team-building");
    name_to_slug.insert("ふりかえり", "reflection");
    name_to_slug.insert("プロジェクトマネジメント", "project-management-advanced");
    name_to_slug.insert("AI", "ai");
    name_to_slug.insert("ビジネス", "business");
    name_to_slug.insert("生産性向上", "productivity-improvement");
    name_to_slug.insert("Google", "google");
    name_to_slug.insert("マーケティング", "marketing");
    name_to_slug.insert("SEO対策", "seo-strategies");
    name_to_slug.insert("解決", "problem-solving");
    name_to_slug.insert("論理的思考", "logical-thinking");
    name_to_slug.insert("リーダー", "leader");
    name_to_slug.insert("kubernetes", "kubernetes");
    name_to_slug.insert("kubectl", "kubectl");
    name_to_slug.insert("ckad", "ckad");
    name_to_slug.insert("CKA", "cka");
    name_to_slug.insert("プレゼンテーション", "presentation");
    name_to_slug.insert("ロジカルシンキング", "logical-thinking-advanced");
    name_to_slug.insert("Go", "go");
    name_to_slug.insert("dockerfile", "dockerfile");
    name_to_slug.insert("DockerHub", "docker-hub");
    name_to_slug.insert("沼", "quagmire");
    name_to_slug.insert("個人開発", "personal-development");
    name_to_slug.insert("GitHubActions", "github-actions");
    name_to_slug.insert("要件定義", "requirement-definition-advanced");
    name_to_slug.insert("ユースケース", "use-case");
    name_to_slug.insert("ワイヤーフレーム", "wireframe");
    name_to_slug.insert("デザイン設計", "design-planning");
    name_to_slug.insert("Cloud", "cloud");
    name_to_slug.insert("docker-compose", "docker-compose");
    name_to_slug.insert("googlecloud", "googlecloud");
    name_to_slug.insert("cookie", "cookie");
    name_to_slug.insert("Vue.js", "vue-js");
    name_to_slug.insert("Vuex", "vuex");
    name_to_slug.insert("ssr", "ssr");
    name_to_slug.insert("開発環境", "development-environment");
    name_to_slug.insert("TypeScript", "typescript");
    name_to_slug.insert("Nuxt3", "nuxt3");

    name_to_slug
}
