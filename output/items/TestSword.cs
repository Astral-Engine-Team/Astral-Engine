using Terraria.ID;
using Terraria.ModLoader;
using Terraria.GameContent.Creative;

namespace ExampleMod.Content.Items
{
    public class TestSword : ModItem
    {
        public override void SetStaticDefaults()
        {
            CreativeItemSacrificesCatalog.Instance.SacrificeCountNeededByItemId[Type] = 100;
        }

        public override void SetDefaults()
        {
            Item.width = 40;
            Item.height = 40;
            Item.maxStack = 1;
            Item.value = 1000;
            Item.rare = ItemRarityID.Green;
        }
    }
}
