## ミーティングのゴール
GitHub Copilotの使い方および使い所を理解している状態

## アジェンダ
1. イントロダクション
  - チェックイン
  - [公式](https://github.com/features/copilot)
  - [基本的な使い方](https://modern-times.dev/github-copilot-for-beginners#784d95309c1d41b1a329d49c5a3804a8)
2. GitHub Copilotのインストール
  - [VS Code](https://docs.github.com/en/copilot/getting-started-with-github-copilot?tool=vscode#installing-the-visual-studio-code-extension)
  - [JetBrains IDE](https://docs.github.com/en/copilot/getting-started-with-github-copilot?tool=jetbrains#installing-the-github-copilot-extension-in-your-jetbrains-ide)
  - [Vim/Neovim](https://docs.github.com/en/copilot/getting-started-with-github-copilot?tool=vimneovim#installing-the-vimneovim-extension-on-macos)
  - [Xcode](https://github.com/intitni/CopilotForXcode#installation-and-setup)
3. GitHub Copilotを使ってみる
  - 提案の表示
    - 主なショートカット
      - 提案の表示: Option (⌥)+\
      - 次の提案の表示: Option (⌥)+]
      - 提案の受け入れ: Tab
      - 提案の拒否: Esc
    - 繰り返し書く処理をCopilotに書いてもらおう
      - Goの場合は `src/error.go` を開いてください
      - Rubyの場合は `src/initialize.rb` を開いてください
    - コード上のコメントをCopilotに書いてもらおう
      - Goの場合は `src/fizzbuzz.go` を開いてください
      - Rubyの場合は `src/fizzbuzz.rb` を開いてください
    - コミットメッセージをCopilotに書いてもらおう
      - コミットメッセージをエディターで書いている方があんまりいなかったらスキップ
  - コメントからコード候補を生成
    - コメントから新規の関数をCopilotに書いてもらおう
      - Goの場合は `src/person.go` を開いてください
      - Rubyの場合は `src/person.rb` を開いてください
    - 既存の関数の説明をCopilotにしてもらおう
      - Goの場合は `src/fizzbuzz.go` を開いてください
      - Rubyの場合は `src/fizzbuzz.rb` を開いてください
4. 質疑応答、ラップアップ

## 備考
- GitHub Copilot は、多数の言語とさまざまなフレームワークに対する候補を提示しますが、特に Python、JavaScript、TypeScript、Ruby、Go、C#、C++ に適しています。
- Given public sources are predominantly in English, GitHub Copilot will likely work less well in scenarios where natural language prompts provided by the developer are not in English and/or are grammatically incorrect. Therefore, non-English speakers might experience a lower quality of service.
