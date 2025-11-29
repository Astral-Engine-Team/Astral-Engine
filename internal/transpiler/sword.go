package transpiler

import (
	"fmt"
	"strings"

	"astral/internal/model"
)

func SwordToCS(s model.Sword) string {
	var b strings.Builder

	b.WriteString("using Terraria.ID;\n")
	b.WriteString("using Terraria.ModLoader;\n")
	b.WriteString("using Terraria.GameContent.Creative;\n\n")

	b.WriteString(fmt.Sprintf("namespace %s\n", s.Namespace))
	b.WriteString("{\n")
	b.WriteString(fmt.Sprintf("    public class %s : ModItem\n", s.Name))
	b.WriteString("    {\n")
	b.WriteString("        public override void SetStaticDefaults()\n")
	b.WriteString("        {\n")
	b.WriteString("            CreativeItemSacrificesCatalog.Instance.SacrificeCountNeededByItemId[Type] = 100;\n")
	b.WriteString("        }\n\n")

	b.WriteString("        public override void SetDefaults()\n")
	b.WriteString("        {\n")
	b.WriteString(fmt.Sprintf("            Item.width = %d;\n", s.Width))
	b.WriteString(fmt.Sprintf("            Item.height = %d;\n", s.Height))
	b.WriteString(fmt.Sprintf("            Item.maxStack = %d;\n", s.MaxStack))
	b.WriteString(fmt.Sprintf("            Item.value = %d;\n", s.Value))
	b.WriteString(fmt.Sprintf("            Item.rare = ItemRarityID.%s;\n", s.Rarity))
	b.WriteString("        }\n")
	b.WriteString("    }\n")
	b.WriteString("}\n")

	return b.String()
}
