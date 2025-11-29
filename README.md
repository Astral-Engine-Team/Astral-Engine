# Astral Engine 🚀

**Write mods in JSON, transpile to native code.** A universal modding engine that lets you create game content with simple, declarative JSON files—starting with Terraria.

✨ **No C# knowledge required** for basic items
⚡ **90% less code** than traditional modding  
🎯 **One syntax, multiple games** (Terraria first, more coming)
🔧 **Fully extensible** with custom logic when needed

If you are looking to include this onto your game? go ahead. have fun!

turn
{
  "name": "CosmicSword",
  "displayName": "Sword of the Cosmos", 
  "damage": 125,
  "rarity": "Lime",
  "useStyle": "Swing"
}

into
using Terraria;
using Terraria.ModLoader;

namespace AstralMod.Content.Items
{
    public class CosmicSword : ModItem
    {
        public override void SetStaticDefaults() { /* ... */ }
        public override void SetDefaults() 
        {
            Item.damage = 125;
            Item.rare = ItemRarityID.Lime;
            Item.useStyle = ItemUseStyleID.Swing;
            // ... and all the boilerplate
        }
    }
}

IN SECONDS!
