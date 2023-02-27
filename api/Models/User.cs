using System.ComponentModel.DataAnnotations.Schema;

namespace pastemyst.Models;

public class User
{
    public string Id { get; set; } = null!;
    
    public DateTime CreatedAt { get; set; }

    [Column(TypeName = "citext")]
    public string Username { get; set; } = null!;
}
