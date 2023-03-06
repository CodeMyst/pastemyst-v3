using System.Net;
using Microsoft.EntityFrameworkCore;
using pastemyst.DbContexts;
using pastemyst.Exceptions;
using pastemyst.Models;
using pastemyst.Utils;

namespace pastemyst.Services;

public interface IPasteService
{
    public Task<Paste> CreatePasteAsync(PasteCreateInfo createInfo);

    public Task<bool> ExistsByIdAsync(string id);
}

public class PasteService : IPasteService
{
    private readonly IIdProvider _idProvider;
    private readonly ILanguageProvider _languageProvider;
    private readonly IPastyService _pastyService;
    private readonly IAuthService _authService;
    private readonly DataContext _dbContext;
    private readonly IHttpContextAccessor _contextAccessor;

    public PasteService(IIdProvider idProvider, DataContext dbContext, ILanguageProvider languageProvider,
        IPastyService pastyService, IAuthService authService, IHttpContextAccessor contextAccessor)
    {
        _idProvider = idProvider;
        _dbContext = dbContext;
        _languageProvider = languageProvider;
        _pastyService = pastyService;
        _authService = authService;
        _contextAccessor = contextAccessor;
    }

    public async Task<Paste> CreatePasteAsync(PasteCreateInfo createInfo)
    {
        var user = await _authService.GetSelfAsync(_contextAccessor.HttpContext);

        if (createInfo.Private && user is null)
        {
            throw new HttpException(HttpStatusCode.Unauthorized,
                "Can't create a private paste while unauthorized.");
        }

        if (createInfo.Private && createInfo.Anonymous)
        {
            throw new HttpException(HttpStatusCode.BadRequest, "Can't create a private anonymous paste.");
        }

        var paste = new Paste
        {
            Id = await _idProvider.GenerateId(async id => await ExistsByIdAsync(id)),
            CreatedAt = DateTime.UtcNow,
            ExpiresIn = createInfo.ExpiresIn,
            DeletesAt = ExpiresInUtils.ToDeletesAt(DateTime.UtcNow, createInfo.ExpiresIn),
            Title = createInfo.Title,
            Owner = createInfo.Anonymous ? null : user,
            Private = createInfo.Private,
            Pasties = new List<Pasty>()
        };

        foreach (var pasty in createInfo.Pasties)
        {
            var langName = pasty.Language is null ? "Text" : _languageProvider.FindByName(pasty.Language).Name;

            paste.Pasties.Add(new Pasty
            {
                Id = await _idProvider.GenerateId(async id => await _pastyService.ExistsByIdAsync(id)),
                Tile = pasty.Title,
                Content = pasty.Content,
                Language = langName
            });
        }

        await _dbContext.Pastes.AddAsync(paste);
        await _dbContext.SaveChangesAsync();

        return paste;
    }

    public async Task<bool> ExistsByIdAsync(string id)
    {
        return await _dbContext.Pastes.AnyAsync(p => p.Id == id);
    }
}