using Microsoft.AspNetCore.Mvc;
using pastemyst.Models;
using pastemyst.Services;

namespace pastemyst.Controllers;

[ApiController]
[Route("/api/v3/pastes")]
public class PasteController : ControllerBase
{
    private readonly IPasteService _pasteService;

    public PasteController(IPasteService pasteService)
    {
        _pasteService = pasteService;
    }

    [HttpPost]
    public async Task<IActionResult> CreatePaste([FromBody] PasteCreateInfo createInfo)
    {
        var paste = await _pasteService.CreatePasteAsync(createInfo);
        return Created("/api/v3/pastes/" + paste.Id, paste);
    }
}