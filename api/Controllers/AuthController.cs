using System.Net;
using Microsoft.AspNetCore.Mvc;
using pastemyst.Exceptions;
using pastemyst.Extensions;
using pastemyst.Models;
using pastemyst.Models.Auth;
using pastemyst.Services;

namespace pastemyst.Controllers;

[ApiController]
[Route("/api/v3")]
public class AuthController(AuthService authService, UserContext userContext) : ControllerBase
{
    [HttpGet("login/{provider}")]
    public async Task<IActionResult> Login(string provider)
    {
        return Redirect(await authService.InitiateLoginFlowAsync(provider, HttpContext));
    }

    [HttpGet("login/{provider}/callback")]
    public async Task<IActionResult> HandleCallback(string provider, [FromQuery] string state, [FromQuery] string code)
    {
        return Redirect(await authService.HandleCallbackAsync(provider, state, code, HttpContext));
    }

    [HttpPost("auth/register")]
    public async Task<IActionResult> Register([FromBody] RegisterRequest registerRequest)
    {
        await authService.RegisterUserAsync(registerRequest.Username, HttpContext);
        return Ok();
    }

    [HttpGet("auth/self")]
    public User GetSelf()
    {
        if (!userContext.IsLoggedIn())
        {
            throw new HttpException(HttpStatusCode.Unauthorized, "You must be authorized to get self.");
        }

        if (!userContext.HasScope(Scope.User, Scope.UserRead))
        {
            throw new HttpException(HttpStatusCode.Forbidden, $"Missing required scope {Scope.UserRead.ToEnumString()}.");
        }

        return userContext.Self;
    }

    [HttpGet("auth/logout")]
    public IActionResult Logout()
    {
        return Redirect(authService.Logout(HttpContext));
    }
}
