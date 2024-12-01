document.addEventListener("turbo:load", () => {
  const selectElement = document.getElementById("user_experience_select");
  const descriptionElement = document.getElementById("experience_description");

  // 選択肢ごとの説明データ
  const descriptions = {
    1: "選択してください",
    2: "これからプログラミングを学び始めたいもしくは学び始めたばかりの方",
    3: "基本的な文法や簡単なプログラムが理解できる方",
    4: "小規模なプロジェクトを自力で構築できる方",
    5: "複数人での開発経験があり、フレームワークを使いこなせる方"
  };

  // 初期表示（選択されているオプションがあればその説明を表示）
  const initialValue = selectElement.value || 1; // 初期選択値がない場合は1をデフォルトに
  descriptionElement.textContent = descriptions[initialValue];

  // 選択変更時の処理
  selectElement.addEventListener("change", (event) => {
    const selectedValue = event.target.value;
    descriptionElement.textContent = descriptions[selectedValue] || "選択肢の説明がありません";
  });
});