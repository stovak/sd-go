// stable_diffusion_api.h

#ifndef STABLE_DIFFUSION_API_H
#define STABLE_DIFFUSION_API_H

#ifdef __cplusplus
extern "C" {
#endif

#include <stdint.h>
#include <stddef.h>

// Opaque handle to the Stable Diffusion context
typedef void* sd_context_t;

// Load the Stable Diffusion model and return a context handle
sd_context_t sd_load_model(const char* model_path);

// Free the Stable Diffusion context
void sd_free_model(sd_context_t ctx);

// Encode input images into latent representations
// - images: pointer to input image data (float32 array)
// - batch_size: number of images in the batch
// - channels: number of channels per image (e.g., 3 for RGB)
// - height: height of each image
// - width: width of each image
// - latents_out: pointer to output buffer for latents (float32 array)
// Returns 0 on success, non-zero on failure
int sd_encode_latents(sd_context_t ctx, const float* images, size_t batch_size, size_t channels, size_t height, size_t width, float* latents_out);

// Perform a forward pass through the model
// - latents: pointer to input latent data (float32 array)
// - batch_size: number of latent samples
// - channels: number of channels per latent
// - height: height of each latent
// - width: width of each latent
// - timestep: diffusion timestep
// - conditioning: pointer to conditioning data (e.g., text embeddings)
// - conditioning_size: size of the conditioning data
// - output: pointer to output buffer (float32 array)
// Returns 0 on success, non-zero on failure
int sd_forward(sd_context_t ctx, const float* latents, size_t batch_size, size_t channels, size_t height, size_t width, int timestep, const float* conditioning, size_t conditioning_size, float* output);

// Compute gradients by comparing predicted output to target
// - target: pointer to target data (float32 array)
// - prediction: pointer to predicted data (float32 array)
// - gradient_out: pointer to output buffer for gradients (float32 array)
// - size: number of elements in the arrays
// Returns 0 on success, non-zero on failure
int sd_backward(sd_context_t ctx, const float* target, const float* prediction, float* gradient_out, size_t size);

// Perform an optimizer step using the computed gradients
// - gradients: pointer to gradient data (float32 array)
// - size: number of elements in the gradient array
// Returns 0 on success, non-zero on failure
int sd_step(sd_context_t ctx, const float* gradients, size_t size);

// Save the current model state to a checkpoint file
// - path: file path to save the checkpoint
// Returns 0 on success, non-zero on failure
int sd_save_checkpoint(sd_context_t ctx, const char* path);

#ifdef __cplusplus
}
#endif

#endif // STABLE_DIFFUSION_API_H
